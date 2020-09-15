// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes a predictor created using the CreatePredictor () operation. In
// addition to listing the properties provided in the CreatePredictor request, this
// operation lists the following properties:
//
//     * DatasetImportJobArns - The
// dataset import jobs used to import training data.
//
//     * AutoMLAlgorithmArns -
// If AutoML is performed, the algorithms that were evaluated.
//
//     *
// CreationTime
//
//     * LastModificationTime
//
//     * Status
//
//     * Message - If an
// error occurred, information about the error.
func (c *Client) DescribePredictor(ctx context.Context, params *DescribePredictorInput, optFns ...func(*Options)) (*DescribePredictorOutput, error) {
	stack := middleware.NewStack("DescribePredictor", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribePredictorMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDescribePredictorValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePredictor(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "DescribePredictor",
			Err:           err,
		}
	}
	out := result.(*DescribePredictorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePredictorInput struct {
	// The Amazon Resource Name (ARN) of the predictor that you want information about.
	PredictorArn *string
}

type DescribePredictorOutput struct {
	// If an error occurred, an informational message about the error.
	Message *string
	// When PerformAutoML is specified, the ARN of the chosen algorithm.
	AutoMLAlgorithmArns []*string
	// The status of the predictor. States include:
	//
	//     * ACTIVE
	//
	//     *
	// CREATE_PENDING, CREATE_IN_PROGRESS, CREATE_FAILED
	//
	//     * DELETE_PENDING,
	// DELETE_IN_PROGRESS, DELETE_FAILED
	//
	//     * UPDATE_PENDING, UPDATE_IN_PROGRESS,
	// UPDATE_FAILED
	//
	// The Status of the predictor must be ACTIVE before you can use the
	// predictor to create a forecast.
	Status *string
	// An AWS Key Management Service (KMS) key and the AWS Identity and Access
	// Management (IAM) role that Amazon Forecast can assume to access the key.
	EncryptionConfig *types.EncryptionConfig
	// The Amazon Resource Name (ARN) of the algorithm used for model training.
	AlgorithmArn *string
	// Describes the dataset group that contains the data to use to train the
	// predictor.
	InputDataConfig *types.InputDataConfig
	// Details on the the status and results of the backtests performed to evaluate the
	// accuracy of the predictor. You specify the number of backtests to perform when
	// you call the operation.
	PredictorExecutionDetails *types.PredictorExecutionDetails
	// The name of the predictor.
	PredictorName *string
	// The featurization configuration.
	FeaturizationConfig *types.FeaturizationConfig
	// An array of the ARNs of the dataset import jobs used to import training data for
	// the predictor.
	DatasetImportJobArns []*string
	// The default training parameters or overrides selected during model training. If
	// using the AutoML algorithm or if HPO is turned on while using the DeepAR+
	// algorithms, the optimized values for the chosen hyperparameters are returned.
	// For more information, see aws-forecast-choosing-recipes ().
	TrainingParameters map[string]*string
	// The number of time-steps of the forecast. The forecast horizon is also called
	// the prediction length.
	ForecastHorizon *int32
	// The hyperparameter override values for the algorithm.
	HPOConfig *types.HyperParameterTuningJobConfig
	// The ARN of the predictor.
	PredictorArn *string
	// Whether the predictor is set to perform AutoML.
	PerformAutoML *bool
	// When the model training task was created.
	CreationTime *time.Time
	// Initially, the same as CreationTime (when the status is CREATE_PENDING). This
	// value is updated when training starts (when the status changes to
	// CREATE_IN_PROGRESS), and when training has completed (when the status changes to
	// ACTIVE) or fails (when the status changes to CREATE_FAILED).
	LastModificationTime *time.Time
	// Whether the predictor is set to perform hyperparameter optimization (HPO).
	PerformHPO *bool
	// Used to override the default evaluation parameters of the specified algorithm.
	// Amazon Forecast evaluates a predictor by splitting a dataset into training data
	// and testing data. The evaluation parameters define how to perform the split and
	// the number of iterations.
	EvaluationParameters *types.EvaluationParameters

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribePredictorMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePredictor{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePredictor{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribePredictor(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "DescribePredictor",
	}
}