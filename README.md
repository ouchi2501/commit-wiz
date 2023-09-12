# commit-wiz

commit-wiz is a utility that summarizes the differences in a Git repository and automatically generates commit comments using OpenAI.

```shell
$ ./build/commit-wiz
Loading... -
The generated commit message is below:

Add commit-wiz utility to automatically generate commit comments using OpenAI

```

## Features

- Retrieves and summarizes Git repository differences.
- Customizable summary length.
- Automatically generates commit comments using the OpenAI API.

## Prerequisites

- Git must be installed.
- OpenAI API key (available from [OpenAI](https://beta.openai.com/signup/)).

## Installation and Setup

1. Clone this repository.

```shell
git clone https://github.com/ouchi2501/commit-wiz.git
```

1. Navigate to the project directory.
```shell
cd commit-wiz
```
1. Install the required Go dependencies.
```shell
go get
```

1. Build the tool.
```shell
go build
```
This will generate an executable binary in the project directory.  
If you're using an ARM-based Mac, you can find the binary in the build directory.

## Usage
Run the tool to summarize Git differences and generate commit comments.
```shell
./build/commit-wiz -length 50
```

Replace 50 with the desired summary length in tokens.

## Configuration
Set up your OpenAI API key by exporting it as an environment variable:
```shell
export OPENAI_API_KEY=your-api-key-here
```

## License
This project is licensed under the MIT License.
See the [LICENSE](https://github.com/ouchi2501/commit-wiz/blob/main/LICENSE) file for details
