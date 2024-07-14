package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

// read this : https://docs.aws.amazon.com/cdk/v2/guide/work-with-cdk-go.html

type GoCdkStackProps struct {
	awscdk.StackProps
}

func NewGoCdkStack(scope constructs.Construct, id string, props *GoCdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	table := awsdynamodb.NewTable(stack, jsii.String("myUserTable"), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("username"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		TableName: jsii.String("userTable"),
	})

	// The code that defines your stack goes here
	myFunc := awslambda.NewFunction(stack, jsii.String("myLambdaFunction"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_PROVIDED_AL2023(),                                    // determines the runtime like Node, Go etc
		Code:    awslambda.AssetCode_FromAsset(jsii.String("lambda/function.zip"), nil), // jsii.String basically converts Go string to ts string, here we are defining as to where is the lambda code written
		Handler: jsii.String("main"),
	})

	table.GrantReadWriteData(myFunc)

	// Now before starting with the deployment process and building the
	// function.zip file inside the lambda module, we need to add the following before the
	// go.build command: 'GOOS=linux GOARCH=amd64 go build -o bootstrap'
	// 'zip function.zip bootstrap' this is used to create the .zip file as specified above

	return stack
}

func main() {
	defer jsii.Close() // built by AWS CDK to run a non TS code
	// defer tells Go to run this afer everything has been ran

	app := awscdk.NewApp(nil)

	NewGoCdkStack(app, "GoCdkStack", &GoCdkStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return nil
}
