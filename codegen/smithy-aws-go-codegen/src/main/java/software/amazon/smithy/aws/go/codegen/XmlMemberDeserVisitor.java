package software.amazon.smithy.aws.go.codegen;

import software.amazon.smithy.codegen.core.CodegenException;
import software.amazon.smithy.codegen.core.Symbol;
import software.amazon.smithy.go.codegen.CodegenUtils;
import software.amazon.smithy.go.codegen.GoWriter;
import software.amazon.smithy.go.codegen.SmithyGoDependency;
import software.amazon.smithy.go.codegen.integration.ProtocolGenerator;
import software.amazon.smithy.go.codegen.integration.ProtocolGenerator.GenerationContext;
import software.amazon.smithy.model.shapes.BigDecimalShape;
import software.amazon.smithy.model.shapes.BigIntegerShape;
import software.amazon.smithy.model.shapes.BlobShape;
import software.amazon.smithy.model.shapes.BooleanShape;
import software.amazon.smithy.model.shapes.ByteShape;
import software.amazon.smithy.model.shapes.CollectionShape;
import software.amazon.smithy.model.shapes.DocumentShape;
import software.amazon.smithy.model.shapes.DoubleShape;
import software.amazon.smithy.model.shapes.FloatShape;
import software.amazon.smithy.model.shapes.IntegerShape;
import software.amazon.smithy.model.shapes.ListShape;
import software.amazon.smithy.model.shapes.LongShape;
import software.amazon.smithy.model.shapes.MapShape;
import software.amazon.smithy.model.shapes.MemberShape;
import software.amazon.smithy.model.shapes.OperationShape;
import software.amazon.smithy.model.shapes.ResourceShape;
import software.amazon.smithy.model.shapes.ServiceShape;
import software.amazon.smithy.model.shapes.SetShape;
import software.amazon.smithy.model.shapes.Shape;
import software.amazon.smithy.model.shapes.ShapeVisitor;
import software.amazon.smithy.model.shapes.ShortShape;
import software.amazon.smithy.model.shapes.StringShape;
import software.amazon.smithy.model.shapes.StructureShape;
import software.amazon.smithy.model.shapes.TimestampShape;
import software.amazon.smithy.model.shapes.UnionShape;
import software.amazon.smithy.model.traits.EnumTrait;
import software.amazon.smithy.model.traits.TimestampFormatTrait.Format;

/**
 * Visitor to generate member values for aggregate types deserialized from documents.
 */
public class XmlMemberDeserVisitor implements ShapeVisitor<Void> {
    private final GenerationContext context;
    private final String dataDest;
    private final Format timestampFormat;

    // isXmlAttributeMember indicates if member is deserialized from the xml start elements attribute value.
    private final boolean isXmlAttributeMember;
    private final boolean isFlattened;

    public XmlMemberDeserVisitor(
            GenerationContext context, String dataDest, Format timestampFormat,
            boolean isXmlAttributeMember, boolean isFlattened
    ) {
        this.context = context;
        this.dataDest = dataDest;
        this.timestampFormat = timestampFormat;
        this.isXmlAttributeMember = isXmlAttributeMember;
        this.isFlattened = isFlattened;
    }

    @Override
    public Void blobShape(BlobShape shape) {
        GoWriter writer = context.getWriter();
        writer.write("var data string");
        handleString(shape, () -> writer.write("data = xtv"));

        writer.addUseImports(SmithyGoDependency.BASE64);
        writer.write("$L, err = base64.StdEncoding.DecodeString(data)", dataDest);
        writer.write("if err != nil { return err }");
        return null;
    }

    @Override
    public Void booleanShape(BooleanShape shape) {
        GoWriter writer = context.getWriter();
        writer.addUseImports(SmithyGoDependency.FMT);
        consumeToken(shape);

        writer.openBlock("if val != nil {", "}", () -> {
            writer.addUseImports(SmithyGoDependency.STRCONV);
            writer.write("xtv, err := strconv.ParseBool(string(val))");
            writer.openBlock("if err != nil {", "}", () -> {
                writer.write("return fmt.Errorf(\"expected $L to be of type *bool, got %T instead\", val)",
                        shape.getId().getName());
            });
            writer.write("$L = &xtv", dataDest);
        });
        return null;
    }

    /**
     * Consumes a single token into the variable "val", returning on any error.
     * If member is an xmlAttributeMember, "attr" representing xml attribute value is in scope.
     */
    private void consumeToken(Shape shape) {
        GoWriter writer = context.getWriter();
        // if the member is a modeled as an xml attribute, we do not need to
        // get another token, instead use the attribute values from previously
        // decoded start element.
        if (isXmlAttributeMember) {
            writer.write("val := []byte(attr.Value)");
            return;
        }

        writer.write("val, done, err := decoder.Value()");
        writer.write("if err != nil { return err }");
        writer.openBlock("if done {", "}", () -> {
            handleDone(shape);
            writer.write("break");
        });
    }

    // handles zero value assignment in case a empty or self closed xml tag is sent for the response member
    private void handleDone(Shape shape) {
        GoWriter writer = context.getWriter();

        // Member shapes string, blob, collection and map can have an empty xml or self closed tag in response.
        // The string, blob shape with empty xml or self closed tag in response are deserialized as empty string.
        // Collection, map shape are deserialized as empty list, map respectively.
        //
        // empty values are handled for collection, map in their respective member deser func.
        //
        // Handle string shape zero value assignment
        if ((shape.isStringShape() && !shape.hasTrait(EnumTrait.ID))) {
            writer.addUseImports(SmithyGoDependency.SMITHY_PTR);
            // assign empty string as zero value if val for string member is nil.
            writer.write("if val == nil { $L = ptr.String(\"\")}", dataDest);
        }

        // Handle blob shape zero value assignment
        if (shape.isBlobShape()) {
            writer.addUseImports(SmithyGoDependency.SMITHY_PTR);
            // assign empty byte slice as zero value if val for blob member is nil.
            writer.write("if val == nil { $L = []byte{}}", dataDest);
        }
    }

    @Override
    public Void byteShape(ByteShape shape) {
        GoWriter writer = context.getWriter();
        // Smithy's byte shape represents a signed 8-bit int, which doesn't line up with Go's unsigned byte
        handleInteger(shape, CodegenUtils.generatePointerValueIfPointable(writer, shape, "int8(i64)"));
        return null;
    }

    @Override
    public Void shortShape(ShortShape shape) {
        GoWriter writer = context.getWriter();
        handleInteger(shape, CodegenUtils.generatePointerValueIfPointable(writer, shape, "int16(i64)"));
        return null;
    }

    @Override
    public Void integerShape(IntegerShape shape) {
        GoWriter writer = context.getWriter();
        handleInteger(shape, CodegenUtils.generatePointerValueIfPointable(writer, shape, "int32(i64)"));
        return null;
    }

    @Override
    public Void longShape(LongShape shape) {
        handleInteger(shape, "&i64");
        return null;
    }

    /**
     * Deserializes a string representing number without a fractional value.
     * The 64-bit integer representation of the number is stored in the variable {@code i64}.
     *
     * @param shape The shape being deserialized.
     * @param cast  A wrapping of {@code i64} to cast it to the proper type.
     */
    private void handleInteger(Shape shape, String cast) {
        GoWriter writer = context.getWriter();
        handleNumber(shape, () -> {
            writer.addUseImports(SmithyGoDependency.STRCONV);
            writer.write("i64, err := strconv.ParseInt(xtv, 10, 64)");
            writer.write("if err != nil { return err }");
            writer.write("$L = $L", dataDest, cast);
        });
    }

    /**
     * Deserializes a xml number string into a xml token.
     * The number token is stored under the variable {@code xtv}.
     *
     * @param shape The shape being deserialized.
     * @param r     A runnable that runs after the value has been parsed, before the scope closes.
     */
    private void handleNumber(Shape shape, Runnable r) {
        GoWriter writer = context.getWriter();
        writer.addUseImports(SmithyGoDependency.FMT);
        consumeToken(shape);

        writer.openBlock("if val != nil {", "}", () -> {
            writer.write("xtv := string(val)");
            r.run();
        });
    }

    @Override
    public Void floatShape(FloatShape shape) {
        GoWriter writer = context.getWriter();
        handleFloat(shape, CodegenUtils.generatePointerValueIfPointable(writer, shape, "float32(f64)"));
        return null;
    }

    @Override
    public Void doubleShape(DoubleShape shape) {
        handleFloat(shape, "&f64");
        return null;
    }

    /**
     * Deserializes a string representing number with a fractional value.
     * The 64-bit float representation of the number is stored in the variable {@code f64}.
     *
     * @param shape The shape being deserialized.
     * @param cast  A wrapping of {@code f64} to cast it to the proper type.
     */
    private void handleFloat(Shape shape, String cast) {
        GoWriter writer = context.getWriter();
        handleNumber(shape, () -> {
            writer.write("f64, err := strconv.ParseFloat(xtv, 64)");
            writer.write("if err != nil { return err }");
            writer.write("$L = $L", dataDest, cast);
        });
    }

    @Override
    public Void stringShape(StringShape shape) {
        GoWriter writer = context.getWriter();
        Symbol symbol = context.getSymbolProvider().toSymbol(shape);

        if (shape.hasTrait(EnumTrait.class)) {
            handleString(shape, () -> writer.write("$L = $P(xtv)", dataDest, symbol));
        } else {
            handleString(shape, () -> writer.write("$L = &xtv", dataDest));
        }

        return null;
    }

    /**
     * Deserializes a xml string into a xml token.
     * The number token is stored under the variable {@code xtv}.
     *
     * @param shape The shape being deserialized.
     * @param r     A runnable that runs after the value has been parsed, before the scope closes.
     */
    private void handleString(Shape shape, Runnable r) {
        GoWriter writer = context.getWriter();
        writer.addUseImports(SmithyGoDependency.FMT);
        consumeToken(shape);

        writer.openBlock("if val != nil {", "}", () -> {
            writer.write("xtv := string(val)");
            r.run();
        });
    }

    @Override
    public Void timestampShape(TimestampShape shape) {
        GoWriter writer = context.getWriter();
        writer.addUseImports(SmithyGoDependency.SMITHY_TIME);

        switch (timestampFormat) {
            case DATE_TIME:
                handleString(shape, () -> {
                    writer.write("t, err := smithytime.ParseDateTime(xtv)");
                    writer.write("if err != nil { return err }");
                    writer.write("$L = &t", dataDest);
                });
                break;
            case HTTP_DATE:
                handleString(shape, () -> {
                    writer.write("t, err := smithytime.ParseHTTPDate(xtv)");
                    writer.write("if err != nil { return err }");
                    writer.write("$L = &t", dataDest);
                });
                break;
            case EPOCH_SECONDS:
                writer.addUseImports(SmithyGoDependency.SMITHY_PTR);
                handleFloat(shape, "ptr.Time(smithytime.ParseEpochSeconds(f64))");
                break;
            default:
                throw new CodegenException(String.format("Unknown timestamp format %s", timestampFormat));
        }
        return null;
    }

    @Override
    public Void bigIntegerShape(BigIntegerShape shape) {
        // Fail instead of losing precision through Number.
        unsupportedShape(shape);
        return null;
    }

    @Override
    public Void bigDecimalShape(BigDecimalShape shape) {
        // Fail instead of losing precision through Number.
        unsupportedShape(shape);
        return null;
    }

    private String unsupportedShape(Shape shape) {
        throw new CodegenException(String.format("Cannot deserialize shape type %s on protocol, shape: %s.",
                shape.getType(), shape.getId()));
    }

    @Override
    public Void operationShape(OperationShape shape) {
        throw new CodegenException("Operation shapes cannot be bound to documents.");
    }

    @Override
    public Void resourceShape(ResourceShape shape) {
        throw new CodegenException("Resource shapes cannot be bound to documents.");
    }

    @Override
    public Void serviceShape(ServiceShape shape) {
        throw new CodegenException("Service shapes cannot be bound to documents.");
    }

    @Override
    public Void memberShape(MemberShape shape) {
        throw new CodegenException("Member shapes cannot be bound to documents.");
    }

    @Override
    public Void documentShape(DocumentShape shape) {
        writeDelegateFunction(shape);
        return null;
    }

    @Override
    public Void structureShape(StructureShape shape) {
        writeDelegateFunction(shape);
        return null;
    }

    @Override
    public Void unionShape(UnionShape shape) {
        writeDelegateFunction(shape);
        return null;
    }

    @Override
    public Void listShape(ListShape shape) {
        return collectionShape(shape);
    }

    @Override
    public Void setShape(SetShape shape) {
        return collectionShape(shape);
    }

    private Void collectionShape(CollectionShape shape) {
        if (isFlattened) {
            writeUnwrappedDelegateFunction(shape);
        } else {
            writeDelegateFunction(shape);
        }
        return null;
    }

    @Override
    public Void mapShape(MapShape shape) {
        if (isFlattened) {
            writeUnwrappedDelegateFunction(shape);
        } else {
            writeDelegateFunction(shape);
        }
        return null;
    }

    private void writeDelegateFunction(Shape shape) {
        String functionName = ProtocolGenerator.getDocumentDeserializerFunctionName(shape, context.getProtocolName());
        GoWriter writer = context.getWriter();
        writer.write("nodeDecoder := smithyxml.WrapNodeDecoder(decoder.Decoder, t)");
        writer.openBlock("if err := $L(&$L, nodeDecoder); err != nil {", "}", functionName, dataDest, () -> {
            writer.write("return err");
        });
    }

    private String getUnwrappedDelegateFunctionName(Shape shape) {
        return ProtocolGenerator.getDocumentDeserializerFunctionName(shape, context.getProtocolName()) + "Unwrapped";
    }

    private void writeUnwrappedDelegateFunction(Shape shape) {
        String functionName = getUnwrappedDelegateFunctionName(shape);
        GoWriter writer = context.getWriter();
        writer.write("nodeDecoder := smithyxml.WrapNodeDecoder(decoder.Decoder, t)");
        writer.openBlock("if err := $L(&$L, nodeDecoder); err != nil {", "}", functionName, dataDest, () -> {
            writer.write("return err");
        });
    }
}
