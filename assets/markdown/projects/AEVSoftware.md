# Alset Solar Cybersedan Software

[Documentation](https://aev.zachl.tech)

[Source Code](https://github.com/YamanDevelopment/AEV-Software)

[Solar Car Challenge](https://www.solarcarchallenge.org/challenge/)

## Overview

The Alset Solar Cybersedan Software is a comprehensive full-stack ecosystem developed for the FAUHS AEV solar car, enabling seamless integration and operation of various hardware components such as the Thunderstruck BMS, cameras, GPS module, and custom sound horn. This project was undertaken as part of the Advanced Experimental Vehicles program at FAU High School.

Me alongside a team of ~30 members who actively worked on the car spent an entire year building this car after ~3 years of planning. After I spent a significant time helping with mechanical and electrical portions of the car, Me and the club president decided to make a Tesla style screen for the dashboard to interface with the entire car. This led to me putting together a coding team to help build this project and integrate it into the car before the race started (The race occurred mid July 2024).

More details on the Solar Car Challenge can be found [here](https://www.solarcarchallenge.org/challenge/)

## Development and Collaboration

The project was a collaborative effort within the solar car club's coding subteam.

As the lead frontend developer, I created the cars UI including a comprehensive homepage with speed, BMS, Lap, and Rear camera data; a camera view page with interactive panels for the rear and blindspot cameras; and a debug page for remotely or manually viewing and resetting specific backend services. I also made sure to ensure seamless integration with the backend, also fostering teamwork and communication among team members.

## Key Features

- **Frontend Interface**: Built using Nuxt.js and TailwindCSS, the interface provides a user-friendly experience with a multi-page layout for various functionalities.
- **Realtime Data Display**: The home page showcases a virtual speedometer, rear camera feed, BMS data, and lap timing features which updates live anywhere in the world via a VPN.
- **Camera Integration**: Three cameras (rear and two blind spot) are accessible on a dedicated page, with options for full-screen viewing.
- **Custom Soundboard and Playlist**: A soundboard for custom horn sounds and a music playlist are available on the third page.
- **Settings and Debugging Tools**: Access to settings like WiFi and Bluetooth, along with a debug terminal for developers, is provided.
- **Data Management**: Options to send data to Google Sheets and restart key components (BMS, GPS, VPN) directly from the interface.

## Technical Details

### Languages and Tools

- **Frontend**: Built with Nuxt.js (Vue.js framework) and styled using TailwindCSS. Electron was used to create a desktop application that the Raspberry Pi in the car used.
- **Backend**: Developed in JavaScript, handling data processing and communication with hardware components.
- **3D Visualization**: Utilized ThreeJS for rendering 3D models and animations.
- **Documentation**: Created using Docus, a documentation tool based on Nuxt.js.

### System Architecture

- **Hardware**: The software runs on a Raspberry Pi 5, chosen for its robust performance with 8GB RAM and a quad-core processor.
- **Operating System**: Arch Linux ARM (aarch64) is used for its lightweight and customizable nature.
- **Navigation**: Hyprland, a dynamic tiling Wayland compositor, was used for window management and Aylur's GTK Shell was used as a bar for navigation across the system.
- **Communication Interfaces**: Direct communication with the BMS via USB serial ports, GPS data handled by gpsd, and networking managed by Wireguard VPN to access live car data from the pits or anywhere in the world.
- **Data Storage**: The software and operating system are installed on a microSD card, ensuring portability and ease of updates. All source and necessary system code is available on our Github repository and we had 5 backup SD cards as well as a backup Raspberry Pi for the actual race

## Documentation

The project's documentation is hosted online, providing guidance and technical details for users and developers. Visit the [Alset Solar Cybersedan Documentation](https://aev.zachl.tech) for more information.

## Notes

The Alset Solar Cybersedan Software represents a significant achievement in integrating technology and sustainability, demonstrating the potential of student-led innovation in the field of advanced experimental vehicles. It stands as a testament to teamwork, technical skill, and a passion for engineering.

## Contributors

These were my fellow teammates who were in the coding subteam and worked on the car as well as the software

[Unix/Backend: Thandi Menelas - Str1ke](https://github.com/RealStr1ke)

[Frontend/Backend: Jossaya Camille](https://github.com/jcamille2023)

[Unix/Documentation: Amarnath Patel](https://github.com/jeebuscrossaint)