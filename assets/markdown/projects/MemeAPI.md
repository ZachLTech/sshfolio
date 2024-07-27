# MemeAPI

[Source Code](https://github.com/ZachLTech/RandomMemeRestAPI)

[Website](https://memes.zachl.tech)

[API Endpoint](https://memeapi.zachl.tech)

## Overview

The MemeAPI is a playful yet educational project that I developed to explore the capabilities of building a RESTful API using ElysiaJS and Bun. Initially conceived as a fun way to learn more about REST APIs, this project evolved into an entertaining service providing random memes, which can be accessed via a simple and intuitive API.

Despite being a side project for learning purposes, it incorporates a range of features that showcase a practical implementation of a REST API, complete with file storage using a self-hosted S3 bucket alternative.

## Development and Features

This project allowed me to experiment with ElysiaJS and Bun, utilizing their unique features to deliver an efficient and responsive API. I designed endpoints that can return either JSON or HTML formatted responses, giving users the flexibility to choose how they consume the data. 

The backend leverages Docker for containerized deployment, ensuring easy setup and scalability. On the frontend, I developed a single-page application that acts as both a demonstration of the API and its documentation. The frontend supports mobile devices, making it accessible on the go.

### Key Features

- **Random Meme Endpoints**: Provides random videos or pictures, with an option to return multiple memes at once.
- **Flexible Response Formats**: Users can receive responses in JSON or HTML format based on their preference.
- **Anti-Repetition Mechanism**: Ensures a unique set of memes in multi-meme requests to avoid duplication.
- **Self-Hosted Deployment**: Docker support allows users to deploy the API on their own servers easily.
- **Robust Frontend**: A single-page website with complete documentation and live demo, accessible on both desktop and mobile.

## Technical Details

### Languages and Tools

- **Backend**: Developed using ElysiaJS, running on Bun for enhanced performance.
- **Frontend**: Created with NuxtJS and TailwindCSS for styling & requests.
- **Containerization**: Docker and Docker Compose are used for streamlined deployment and management on my server.

### System Architecture

- **Self-Hosted Storage**: Utilizes a self-hosted S3 bucket alternative for meme storage, providing a scalable and cost-effective solution.
- **Endpoints**: The API features multiple endpoints for different types of content (random, video, picture), with both single and multiple result options.

## Notes

Feel free to explore the [MemeAPI website](https://memes.zachl.tech) to see the API in action and enjoy some memes!
