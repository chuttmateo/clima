# Clima

Clima is a command-line tool written in Golang, utilizing the Cobra framework. It provides weather forecasts for today or the next three days.

## Installation

1. Download the latest version of Clima from [releases](https://github.com/chuttmateo/clima/releases).
2. Add the downloaded binary to your `bin` folder.

## Usage

Before using Clima, you need to set up your WeatherAPI token and optionally specify your location.

### Setting up WeatherAPI token

1. Go to [WeatherAPI](https://www.weatherapi.com/) and create your own token. (for free)
2. Declare an environment variable `CLIMA_TOKEN` and set it to your WeatherAPI token.

### Specifying location (optional)

You can specify your default location by declaring an environment variable `CLIMA_LOCATION`.

## Example

```bash
# Get the current weather
clima current

# Get the forecast for the next 3 days
clima forecast

# Get the forecast for a specific location
clima -l "Las Vegas" current
