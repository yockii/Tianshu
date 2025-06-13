# Tianshu (天枢)
**中文版文档**： [README_zh.md](README_zh.md)

Tianshu is a Platform to manage and control DJI drones. It uses DJI Pilot2 to communicate with the drones.

# Features
- **Drone Management**: Manage multiple drones and their configurations.
- **Flight Planning**: Plan and execute complex flight missions.
- **Real-time Monitoring**: Monitor drone status and telemetry data in real-time.
- **User Management**: Manage user roles and permissions within the platform.
- **Data Logging**: Log flight data for analysis and reporting.
- **API Integration**: Integrate with other systems using RESTful APIs.
- **MQTT Embedding**: Use MQTT for real-time data streaming and communication. You don't need to run a separate MQTT broker.

# Installation
1. Install postgresql and create a database, install redis too.
2. Download the latest release from the release page.
3. Copy config.toml.example to config.toml and edit it.
4. Run the release binary.

# Usage
## For Site Managers
1. Open your web browser and navigate to the Tianshu server URL.
2. Log in with your admin account.

## For smart controllers
1. Connect your smart controller to the drone.
2. Open pilot2 and log in with your DJI account.
3. In 'Cloud Service' / 'Open Platform', enter the URL of your Tianshu server.
4. Click 'Connect' and sign in to link your smart controller with Tianshu.

## Status and Next Steps

- Completed user authentication and Pilot2 login integration.  
- Pilot2 platform login implemented; due to personal financial constraints, industrial-grade hardware is temporarily unavailable, so protocol connection validation is on hold.  
- A virtual Pilot2 environment is under development to support further development and testing (behavior may differ from real hardware).  
- Next steps: finalize the virtual environment and continue full project development; resume real-world testing once hardware becomes available.

We welcome any support or sponsorship to help acquire DJI industrial/enterprise drones for testing. Please stay tuned for our updates!

