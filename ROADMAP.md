# ROADMAP

## v0.9.0-beta

- Add git revision configuration.
- Add deploying from S3 bucket.
    - For this the cf config needs access to the bucket.

## v0.0.1

- Add error test cases for the calls which return errors.
    - For example describeStacks returns an error if the stack is non-existent.
- Add push command which deals with pushing a version of the application to an
existing stack.