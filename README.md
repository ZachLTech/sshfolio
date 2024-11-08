# SSHFolio 🚀

SSH Portfolio is a modular TUI (Text User Interface) portfolio application based on the [Bubble Tea framework](https://github.com/charmbracelet/bubbletea). This project allows you to showcase your portfolio in a terminal environment, providing a unique way to present your projects, biography, contact information, and more. It's designed to be easily customizable and self-hostable using Docker.

## Features 🌟

- **Modular and Customizable**: Easily modify the portfolio to suit your needs using simple `.env` configuration and markdown files for your projects, about me, contact page, and home page.
- **SSH Optional**: Configure the application to run either as an SSH server or directly as a TUI in the terminal where it's executed, depending on your setup preferences.
- **Mouse Support**: 🖱️ Navigate through the application using scrolling and clicking, even though it's a TUI.
- **Self-Hostable**: 🏠 Includes `Dockerfile` and `docker-compose.yml` for easy deployment.
- **Markdown Rendering**: 🎨 Customize markdown colors and behavior via the `/assets/MDStyle.json` file, with support from the [Glamour](https://github.com/charmbracelet/glamour) library.

## SSH Quick Start 🚀

### Prerequisites 📋

- [Docker](https://www.docker.com/get-started) installed on your system.
- An SSH client if you plan to use the SSH server feature.

### Installation 🛠️

1. **Clone the Repository**

   ```bash
   git clone https://github.com/ZachLTech/sshfolio.git
   cd sshfolio
   ```

2. **Set Up Environment Variables**

   Create a `.env` file in the project root and configure it according to your needs. Here's a template to get you started:

   ```env
   SSH_SERVER_ENABLED=true # If false, the application will only serve the TUI in the terminal where the program is being run. If true, the application will run an SSH server for the TUI and serve to the port specified below.
   PORT=23
   HOST="0.0.0.0" # Whatever machine you run this on will be the host by default.

   HEADER="John Doe"
   HEADER_MESSAGE="'The intricacy of John is just so very Doe' - Jane Doe" # Doesn't have to be a quote - but that's what I put for mine :D

   # Project Configuration (Don't leave any extra blank projects. This is just an example)
   PROJECT_1_MARKDOWN_FILE_TITLE="LoremIpsum"
   PROJECT_1_DISPLAY_TITLE="Lorem Ipsum Woah"
   PROJECT_1_DESCRIPTION="Lorem ipsum dolor sit amet, consectetur adipiscing elit"
   ```

3. **Run with Docker** 🐳

   Build and start the Docker container using Docker Compose:

   ```bash
   docker-compose up --build
   ```

4. **SSH** ‼️
   
   Once the docker container has finished building and is running, your SSHFolio instance will be available to SSH into if SSH_SERVER_ENABLED is set to `true`.

   Assuming you're testing this all on the same machine, you can now SSH into the portfolio TUI like so:
   
   ```bash
   ssh localhost -p 23
   ```

### Usage 🎯

- **Make it your own**
  
  If you're planning to use this as your own SSHfolio, here's are a few steps on how to do so.

  1. After cloning, rename the `.env.sample` file to `.env` and fill it out with your own personal portfolio information
  2. Write your Home, About, and Contact page markdown files (they're located in `/assets/markdown`
  3. Write your corresponding Project markdown files for each project you listed in the `.env` file
     - Each markdown file should be named what you put under `PROJECT_x_MARKDOWN_FILE_TITLE` with the .md at the end of course
     - All project markdown files should be placed in the `/assets/markdown/projects` directory
    
  And that's it! You now have your very own, fully customized, portfolio TUI!!

- **Access Your Instance**

  If running as an SSH server, connect to your instance via SSH:

  ```bash
  ssh <HOST-MACHINE-IP> -p <PORT>
  ```

  If running as a local TUI, simply execute the program:

  ```bash
  go mod tidy && go run .
  ```

- **View the Demo**

  Check out the live demo instances:

  - Personal instance: `ssh zachl.tech`
  - Demo instance (Not currently up): ~~`ssh sshdemo.zachl.tech`~~

## Further Customization 🎨

- **Markdown Styles**

  Edit the `/assets/MDStyle.json` file to customize markdown rendering. Refer to the [Glamour styles documentation](https://github.com/charmbracelet/glamour/tree/master/styles) for guidance on available options.

## TODO

- Add an intro and outro screen
- Make the intro and outro screen customizable with .env

## Credits 🙏

SSH Portfolio is built on the amazing [Charm](https://charm.sh/) ecosystem, utilizing the Bubble Tea framework. Huge thanks to Charmbracelet for their incredible work and inspiration. The inspiration for this project was from one of their demos hosted at `ssh git.charm.sh`. And when I saw that you could pair [Bubbletea](https://github.com/charmbracelet/bubbletea) with something they made called [Wish](https://github.com/charmbracelet/wish) to make it an SSHable application, I was absolutely amazed and immediately started working on this project xD

## Contributing 🤝

Contributions are welcome! Please feel free to submit a pull request or open an issue.

## License 📄

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
