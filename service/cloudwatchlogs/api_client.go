// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	"net/http"
)

const ServiceID = "CloudWatch Logs"

// You can use Amazon CloudWatch Logs to monitor, store, and access your log files
// from Amazon EC2 instances, AWS CloudTrail, or other sources. You can then
// retrieve the associated log data from CloudWatch Logs using the CloudWatch
// console, CloudWatch Logs commands in the AWS CLI, CloudWatch Logs API, or
// CloudWatch Logs SDK. You can use CloudWatch Logs to:
//
//     * Monitor logs from
// EC2 instances in real-time: You can use CloudWatch Logs to monitor applications
// and systems using log data. For example, CloudWatch Logs can track the number of
// errors that occur in your application logs and send you a notification whenever
// the rate of errors exceeds a threshold that you specify. CloudWatch Logs uses
// your log data for monitoring; so, no code changes are required. For example, you
// can monitor application logs for specific literal terms (such as
// "NullReferenceException") or count the number of occurrences of a literal term
// at a particular position in log data (such as "404" status codes in an Apache
// access log). When the term you are searching for is found, CloudWatch Logs
// reports the data to a CloudWatch metric that you specify.
//
//     * Monitor AWS
// CloudTrail logged events: You can create alarms in CloudWatch and receive
// notifications of particular API activity as captured by CloudTrail and use the
// notification to perform troubleshooting.
//
//     * Archive log data: You can use
// CloudWatch Logs to store your log data in highly durable storage. You can change
// the log retention setting so that any log events older than this setting are
// automatically deleted. The CloudWatch Logs agent makes it easy to quickly send
// both rotated and non-rotated log data off of a host and into the log service.
// You can then access the raw log data when you need it.
type Client struct {
	options Options
}

// New returns an initialized Client based on the functional options. Provide
// additional functional options to further configure the behavior of the client,
// such as changing the client's endpoint or adding custom middleware behavior.
func New(options Options, optFns ...func(*Options)) *Client {
	options = options.Copy()

	resolveRetryer(&options)

	resolveHTTPClient(&options)

	resolveDefaultEndpointConfiguration(&options)

	for _, fn := range optFns {
		fn(&options)
	}

	client := &Client{
		options: options,
	}

	return client
}

type Options struct {
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []APIOptionFunc

	// The credentials object to use when signing requests.
	Credentials aws.CredentialsProvider

	// The endpoint options to be used when attempting to resolve an endpoint.
	EndpointOptions ResolverOptions

	// The service endpoint resolver.
	EndpointResolver EndpointResolver

	// An integer value representing the logging level.
	LogLevel aws.LogLevel

	// The logger writer interface to write logging messages to.
	Logger aws.Logger

	// The region to send requests to. (Required)
	Region string

	// Retryer guides how HTTP requests should be retried in case of recoverable
	// failures. When nil the API client will use a default retryer.
	Retryer retry.Retryer

	// The HTTP client to invoke API calls with. Defaults to client's default HTTP
	// implementation if nil.
	HTTPClient HTTPClient
}

func (o Options) GetCredentials() aws.CredentialsProvider {
	return o.Credentials
}

func (o Options) GetEndpointOptions() ResolverOptions {
	return o.EndpointOptions
}

func (o Options) GetEndpointResolver() EndpointResolver {
	return o.EndpointResolver
}

func (o Options) GetLogLevel() aws.LogLevel {
	return o.LogLevel
}

func (o Options) GetLogger() aws.Logger {
	return o.Logger
}

func (o Options) GetRegion() string {
	return o.Region
}

func (o Options) GetRetryer() retry.Retryer {
	return o.Retryer
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Copy creates a clone where the APIOptions list is deep copied.
func (o Options) Copy() Options {
	to := o
	to.APIOptions = make([]APIOptionFunc, len(o.APIOptions))
	copy(to.APIOptions, o.APIOptions)
	return to
}

type APIOptionFunc func(*middleware.Stack) error

// NewFromConfig returns a new client from the provided config.
func NewFromConfig(cfg aws.Config, optFns ...func(*Options)) *Client {
	opts := Options{
		Region:      cfg.Region,
		Retryer:     cfg.Retryer,
		LogLevel:    cfg.LogLevel,
		Logger:      cfg.Logger,
		HTTPClient:  cfg.HTTPClient,
		Credentials: cfg.Credentials,
	}
	return New(opts, optFns...)
}

func resolveHTTPClient(o *Options) {
	if o.HTTPClient != nil {
		return
	}
	o.HTTPClient = aws.NewBuildableHTTPClient()
}

func resolveRetryer(o *Options) {
	if o.Retryer != nil {
		return
	}
	o.Retryer = retry.NewStandard()
}

func addClientUserAgent(stack *middleware.Stack) {
	awsmiddleware.AddUserAgentKey("cloudwatchlogs")(stack)
}

func addHTTPSignerV4Middleware(stack *middleware.Stack, o Options) {
	signer := v4.Signer{}
	stack.Finalize.Add(v4.NewSignHTTPRequestMiddleware(o.Credentials, signer), middleware.After)
}