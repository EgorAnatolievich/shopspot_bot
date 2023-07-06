
# Telegram Bot for Selling Goods

![Telegram Logo](Telegram_logo.svg.png)

Telegram Bot for Selling Goods is a project written in Go that includes a web application for selling goods. Payment systems are integrated, and the purchase of goods is implemented through the bot menu button.

[![License](https://img.shields.io/badge/license-CC%20BY--NC--ND%204.0-green.svg)](https://creativecommons.org/licenses/by-nc-nd/4.0/)

## About the Project

This project is a Telegram bot developed in Go language. It is designed to launch a web application for selling goods with integrated payment systems. The bot allows users to browse the product catalog, select and purchase items through a convenient menu.

## Features

- Implementation of a web application for selling goods.
- Integration with payment systems for accepting payments.
- Ability to purchase goods through the bot menu button.

## Installation and Setup

1. Clone the repository on your local machine:
git clone https://github.com/your-username/your-repo.git

2. Install the necessary dependencies:
go mod download


3. Specify the database and payment system settings in the `config.yml` file.

4. Run the web application and the bot:
go run main.go
go run telegrambot.go


5. Open the bot in Telegram and start using its features.

## License

This project is licensed under the [Attribution-NonCommercial-NoDerivatives 4.0 International (CC BY-NC-ND 4.0)](https://creativecommons.org/licenses/by-nc-nd/4.0/) License.
