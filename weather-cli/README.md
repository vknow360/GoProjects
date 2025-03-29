# Weather CLI App

A simple **Command-Line Interface (CLI) Weather App** built with **Go** that fetches real-time weather updates using WeatherAPI.com. Get instant weather details for any city right from your terminal! 🌍🌦️

## Features 🚀

-   🌡️ Get **real-time temperature** (Celsius & Fahrenheit)
-   🥶 **Feels like** temperature support
-   🌤️ Current **weather condition**
-   💨 **Wind speed** and 💧 **humidity** details
-   🔒 Secure **API key handling** with environment variables

## Installation 📥

1. **Clone the repository:**
    ```sh
    git clone https://github.com/yourusername/Weather-CLI.git
    cd Weather-CLI
    ```
2. **Set up your API key:**

    - Get your API key from [WeatherAPI.com](https://www.weatherapi.com/)
    - Set the API key as an environment variable:
        ```sh
        export WEATHER_API_KEY=your_api_key_here
        ```

3. **Run the application:**
    ```sh
    go run main.go <city> [-f]
    ```
    - Replace `<city>` with the desired city name.
    - Add `-f` flag to get the temperature in Fahrenheit.

## Example Usage 📌

```sh
go run main.go London
```

**Output:**

```
🌍 Weather in London, United Kingdom:
🌡️  Temperature: 15.2 °C
🥶 Feels Like: 14.5 °C
🌤️  Condition: Partly Cloudy
💨 Wind Speed: 12.4 km/h
💧 Humidity: 65%
```

## Dependencies 🛠️

-   **[tidwall/gjson](https://github.com/tidwall/gjson)** for efficient JSON parsing
-   **net/http** for API requests
-   **os** for handling environment variables
