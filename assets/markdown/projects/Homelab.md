# The Lopez Lab

## Overview

**The Lopez Lab** is a sophisticated homelab environment that I meticulously designed and built to explore and implement advanced network infrastructure, server management, and virtualization technologies. This setup allows me to host a variety of services, manage extensive data storage, and maintain a robust and secure network.

## Network Security and Structure

To ensure security and efficiency, my network leverages multiple tools and protocols:

- **Nginx Proxy Manager**: Provides a straightforward way to manage reverse proxying and SSL certificates.
- **Cloudflare Tunnel**: Secures external access without extensive port forwarding and more, enhancing security.
- **Automatic SSL Certificates**: Enables automatic SSL certificate issuance and renewal for secure connections.
- **Wireguard VPN**: Allows secure, remote access to my network across all my devices no matter where in the world they are.

## Hardware

My homelab is powered by a combination of robust hardware components:
- **Dell PowerEdge R720 H710 Server**: Equipped with two E5-2690 2.90GHz 16-core processors and 192GB DDR3 RAM.
- **Storage Setup**:
  - 5-drive Yottamaster HDD bay connected via USB 3.0.
  - 2× 4TB mirrored drives and 2× 1TB mirrored drives for redundancy.
  - 1× 500GB unassigned drive for flexible use.
  - 2× 150GB 2.5” HDDs dedicated to Proxmox.
  - 2-port USB 3.0 card for expanded connectivity.
- **Networking**: A basic 5-port switch connects to a Google Home router extender, providing reliable network access.

## Server and Infrastructure Setup

My homelab's architecture is built on Proxmox, utilizing TrueNAS for data management and several virtual machines (VMs) for diverse applications:

- **Proxmox Virtualization**: Utilizes an SMB pool from TrueNAS for VMs, allowing efficient resource allocation and management.
- **TrueNAS Scale**: Implements a ZFS setup with drive mirroring for robust data protection.
- **Dedicated VMs**:
  - **Ubuntu Server**: Specially configured for Nextcloud.
  - **Arch Servers**: Used for Docker services, separated into small storage for public services and large storage for private data.
  - **Appwrite Server**: Supports backend services for personal projects.
  - **Arch Sandbox**: Provides a secure environment for experimentation by myself and friends.
- **Portainer**: Manages Docker environments across servers, supporting over 75 containers currently.
- **Service and Domain Management**: Manages 6+ domains, 40+ public and private services, 5+ self-developed websites, and 3+ REST APIs.
- **Uptime Kuma**: Monitors service availability, providing notifications when downtime occurs to multiple devices of mine.

## Achievements and Capabilities

The Lopez Lab stands as a testament to my dedication to technology and self-hosted solutions. This homelab setup empowers me to:

- Experiment with cutting-edge technologies and improve my technical skills.
- Host and manage a variety of public and private services and websites including ones I make.
- Implement a secure, efficient, and scalable network and server infrastructure.
- Engage in continuous learning and development, adapting to new challenges and opportunities.

### Places my infrastructure is used

- For the AEV Solar Cybersedan dashboard as it connects to my VPN in order to communicate with the pits
- Hosting the backend and frontend for many of my projects such as Devfolify, My MemeAPI, Eduquest , WebDevCourse, Schedulix's frontend, AEV Docs, and more
- Hosts a BeamMP server for me and my friends to play on and occasionally I would host a Minecraft server for friends
- Family members backup and view their photos to my infrastructure as their cloud to avoid google photo or icloud costs and more storage.
- A URL shortenning service under my domain "byeurl.cyou" is hosted and managed by me and my server
- A dropbox alternative which has been put to extensive use in the past when transferring larger files to or from friends/family
- My default search engine on most of my devices is a SearXNG instance being managed by my server

## Notes

Soon I will add to this a list of public services I have and their domains but most of them are private or require authorization.