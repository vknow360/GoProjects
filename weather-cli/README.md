# Weather CLI App 🌦️

A simple **Weather CLI App** built using **Go** that fetches real-time weather updates for any city.

## Features 🚀

-   Get real-time weather updates for any city 🌍
-   Supports both **Celsius & Fahrenheit** 🌡️
-   Displays **"Feels Like" temperature** 🥶
-   Shows wind speed and humidity 💨💧
-   Secure API key handling using environment variables 🔒

## Installation & Build Instructions ⚙️

1. **Clone the repository:**

    ```sh
    git clone https://github.com/yourusername/GoProjects.git
    cd GoProjects/weather-cli
    ```

2. **Set up the API key:**

    - Get a free API key from [WeatherAPI.com](https://www.weatherapi.com/)
    - Set the API key as an environment variable:
        ```sh
        export WEATHER_API_KEY=your_api_key_here
        ```

3. **Build the application:**
    ```sh
    go build weather.go
    ```

## Usage 📌

Run the program with the city name as an argument:

```sh
weather <city>
```

Example:

```sh
weather London
```

To get the temperature in **Fahrenheit**, use the `-f` flag:

```sh
weather London -f
```

## Example Output 🌍

```
🌍 Weather in Gorakhpur, India:
🌡️  Temperature: 36.1 °C
🥶 Feels Like: 34.6 °C
🌤️  Condition: Sunny
💨 Wind Speed: 22.7 km/h
💧 Humidity: 3%
```

## Dependencies 📦

-   Go (1.16+ recommended)
-   [tidwall/gjson](https://github.com/tidwall/gjson) for JSON parsing
