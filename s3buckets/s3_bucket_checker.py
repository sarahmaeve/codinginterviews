import boto3
from botocore.exceptions import ClientError

s3 = boto3.client('s3')
response = s3.list_buckets()

manifest = {}
for bucket in response['Buckets']:
    bucket_name = bucket['Name']
    try:
        get_enc = s3.get_bucket_encryption(Bucket=bucket_name)
        manifest[bucket_name]= get_enc['ServerSideEncryptionConfiguration']['Rules'][0]['ApplyServerSideEncryptionByDefault']['SSEAlgorithm']
    except ClientError as err:
        if err.response['Error']['Code'] in ["ServerSideEncryptionConfigurationNotFoundError"]:
            manifest[bucket_name] = "unencrypted"
            # print("{} has no encryption.".format(bucket_name))

header = "S3 buckets with non-KMS encryption"
prettyline = "-" * (len(header))
print(header, "\n", prettyline, sep='')

for k in manifest:
    if manifest[k] != "aws:kms":
        print("{} : {}".format(k, manifest[k]))