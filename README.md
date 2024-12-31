<p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="https://routehub.link/images/5.jpeg" alt="RouteHUB Dashboard LOGO" title="RouteHUB Dashboard"></a>
</p>

<h3 align="center">RouteHUB.Service.Dashboard</h3>

![rh-dashboard-hub](https://github.com/user-attachments/assets/74510644-38b2-4396-9ed0-69a6e207f252)

<details>

<summary>See the dashboard</summary>

![rh-dashboard-index](https://github.com/user-attachments/assets/42a318c8-266c-4ff4-964a-04ec3bdaf1bd)
![rh-dashboard-attach](https://github.com/user-attachments/assets/55429ddf-d83c-409f-94d3-976e37cf9978)
![rh-dashboard-links](https://github.com/user-attachments/assets/26b87c50-01f7-44f4-92b1-fee67bf228c2)
![rh-dashboard-customize-footer](https://github.com/user-attachments/assets/492c2b3f-aaca-4450-b834-e03dfba79ce6)

</details>

---

<p align="center">
RouteHUB is CMS for Link Shortening and Management. You can deploy and manage multiple shortening services, manage links with SEO and meta tags with crawl feature, create custom pages with grapesjs , implement analytics, and more.
    <br>
</p>

## 📝 Table of Contents

- [About](#-about)
- [Getting Started](#-getting-started)
- [Installation](#installation)
- [Deployment](#-deployment)
- [Usage](#-usage)
- [Authors](#-authors)
- [Acknowledgments](#-acknowledgement)
- [Contributing](#-contributing)

## 🧐 About

RouteHub Service Dashboard is a comprehensive platform designed to manage and customize various aspects of your web applications. It provides a user-friendly interface for managing domains, hubs, links, and pages, along with extensive customization options for meta descriptions, SEO fields, and more.

## 🏁 Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See [deployment](#-deployment) for notes on how to deploy the project on a live system.

### Prerequisites

#### Zitadel is a prerequisite for RouteHub Service Dashboard. You can find the deployment example at [Our Compose](docker/zitadel/zitadel-docker-compose.yaml)

```sh
cd docker/zitadel
docker-compose -f zitadel-docker-compose.yaml up -d
```

#### If you want to take analytics we recommend using [Umami](https://github.com/umami-software/umami) you can deploy the analytics service. You can find the deployment example at [Our Compose](docker/umami/docker-compose.yml)

```sh
cd docker/umami
docker compose up -d --force-recreate
```

#### For deploying Dashboard you need these following requirements:

- NPM & NPX
- Tailwindcss
- GO 1.23
- Docker
- Docker Compose

#### It dosent matter if you deploy the client before or after the server. The client connection will be established when you attach the client to the server via dashboard.

### Installing

A step by step series of examples that tell you how to get a development env running.

## Installation

1. **Clone the repository**:

```sh
 git clone https://github.com/RouteHub-Link/RouteHub.Service.Dashboard.git
 cd RouteHub.Service.Dashboard
```

2. **Install dependencies**:

```sh
 npm install
```

3. **Set up environment variables**:

 Copy the `.copy.env` file to `.env` and update the environment variables as needed.

```sh
 cp .copy.env env
```

4. **Start the Database Instances**:

```sh
 make postgres
 make cache-database
```

5. **Build the project**:

```sh
 make build
```

6. **Run the project**:

```sh
 make run
```

## 🎈 Usage

Open your browser and navigate to access the RouteHub Service Dashboard.
You will be signed in with Zitadel instance.
First of all create a domain this does not have to be a real domain. Creating a tags for Client Hubs.
After the creation of the domain you can Attach the client to the server.
If you did not deploy the client before you can do it now. For deployment you can check the [Client Deployment](https://github.com/RouteHub-Link/RouteHub.Client.Hub).
Or just run those commands:

```sh
# Clone the project
git clone https://github.com/RouteHub-Link/RouteHub.Client.Hub
cd RouteHub.Client.Hub

# Copy .env.development to .env
cp .env.development .env

# Run the services
docker compose -f "docker-compose.hosted.yml" up -d --build
```

Do not forget the env variables in the client.

After those steps you can attach the client to the server. You can see the client in the dashboard.
You need to MQTT connection to the client for that compose you can use like this : 'localhost:1883' or your proxy link.
If the connection is not established you have to make sure service can reach the client mqtt container.

After the connection is established you can see the client in the dashboard. You will redirect to the hub dashboard. You can create a new page, edit home page, create a new link and many more customization options.

## 🚀 Deployment

This are will be updated soon. but you can create a custom docker-compose file for the deployment.
The structure of network like this;

<img src="https://mermaid.ink/svg/pako:eNp9VNuO2jAQ_RXLK1UgJYjrAmlVCQgtSOzCkuwLBFWGDElKEkeOs4UC_17nxiVi6wdr7HNmfDwz9hFvqAlYwRYjgY30vuEjMcJonW4YeE4jDqP3PlIhcOnBA58bOGU9ZEbrysB1BKsizFtmPOZDTf8lgi1jA_Vm49U9_vKm68t4QhqwD2Crors61hL_0hxMJ0Rqv3xDAd80_P9LU0lorylhZlGZ2tNG_Wlvri4vlETtQwVXcqZjQDY2lAvc2VTTf4oYy1JuvU3KKyTLMjppnDLHt5C4i4QGUciph2bEAvQFTRx_h1TCyekq6hOtaaz48DiW2kfT9W_Y8PBU1PlpjhZjvacOJ8uFw4kJLpr23vVRdo8b6vhVH85fh3qpNPY5MB94uZwi93giKC_yA6gg6A7LpKRI0gNyRf5-rXkK5MEfgtfEfEvg0w_gojCiTy41HdjEt0BkKD6h6JU6xSVlcRrRljLkUsvxT_fq0nnjkjBUYYtgH9AQTLR1XFd52nafpZAzugPlqdFoZLb8xzG5rTSD_dcb78t1pIsIKTspj5rRsYQ9YB5xTPFYj_GegbkNHhhYEaZJ2C5u6bPgkYhT7eBvsMJZBBJmNLJsrGyJG4pVFJiEg-oQ8Ta8y25A_AWlXu4illg54j1WqhI-YKVWb1Va7Va30enU29V6rdM-S_hv4lCtdPPRfG5Xm7W6hEUCRX-_pD9L8sGc_wEn7EVO" alt="Network Structure" width="1000px"/>

## ✍️ Authors

- [@runaho](https://github.com/Runaho) - Idea & Initial work

## 🎉 Acknowledgements

A big shoutout to all amazing cod work contributed this project Your creativity and are inspiring. Special thanks to all the sources of that fueled ideas and helped this vision to. also to the valuable references that guided us this journey

## 🤝 Contributing

- Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.
- This project developed as thinking Get Shit Done, Features first beauty later.
- Ask your self everytime is this the easiest way.
- I Recommend you check this document first : (The Grug Brained Developer)[https://grugbrain.dev/]
- Use Standart GO libraries as much as possible.
- Do not send requiest's with too big commits.
