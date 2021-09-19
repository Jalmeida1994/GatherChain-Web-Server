<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/Jalmeida1994/GatherChain-Web-Server">
    <img src="images/Logo-02.png" alt="Logo" width="120">
  </a>

  <h3 align="center">GatherChain Web Server</h3>

  <p align="center">
    Web server for the GatherChain solution.
    <br />
    <a href="https://github.com/Jalmeida1994/GatherChain-Web-Server/blob/master/README.md"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/Jalmeida1994/GatherChain-Web-Server/issues">Report Bug</a>
    ·
    <a href="https://github.com/Jalmeida1994/GatherChain-Web-Server/issues">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

GatherChain is the solution created for my Master Thesis: __Tracing Responsibility in Evolution of Model's Life Cycle in Collaborative Projects in Education__ in the Informatics' Department of __NOVA School of Science and Technology__ made with professor __Vasco Amaral__'s guidance.
In the paper, it is proposed a blockchain-based solution for version control of model-driven engineering artefacts.  The goal is to facilitate collaboration in a multi-user area, like the education field, and track changes in a trusted and secure manner. This solution is based on using the Hyperledger Fabric Network to govern and regulate file version control functions among students and teachers.
This repository is a part of the larger GatherChain solution.

The other GatherChain projects are:
* __GatherChain Desktop Client__: https://github.com/Jalmeida1994/GatherChain-DesktopClient
* __GatherChain ARM Template__: https://github.com/Jalmeida1994/GatherChain-ARM-Template
* __GatherChain Blockchain Server__: https://github.com/Jalmeida1994/GatherChain-BlockChain-Server
* __GatherChain Admin Commands__: https://github.com/Jalmeida1994/GatherChain-AdminCommands

This repository serves as the Web server for the blockchain solution. It acts as a middle-man between the final users and the VM hosting the blockchain network. The Docker image can be found at [jalmeida1994/gatherchain-app](https://hub.docker.com/repository/docker/jalmeida1994/gatherchain-app).

### Built With

The web server is built with:
* [Go](https://golang.org)
* [Docker](https://www.docker.com)


<!-- GETTING STARTED -->
## Getting Started

The web app is used automatically by the [GatherChain ARM Template](https://github.com/Jalmeida1994/GatherChain-ARM-Template). If you want to deploy this web app in your environment
### Prerequisites

If you want to build the Docker image on premises and use it as you own, you'll need:
* [Docker](https://docs.docker.com/get-docker/)

### Installation

If the default installation is followed, using Azure as the cloud provider, you don't need to worry about this repository. The ARM template can be found in the repository [GatherChain ARM Template](https://github.com/Jalmeida1994/GatherChain-ARM-Template).

_For more information, please refer to the [Documentation](https://github.com/Jalmeida1994/GatherChain-ARM-Template/blob/master/README.md)_

If you want to build the Docker image on premises and use it as you own, follow the next steps:


1. Clone the repo
   ```
   git clone https://github.com/Jalmeida1994/GatherChain-Web-Server.git
   ```
2. The server hosting the web app needs to have configured 5 environment variables to work correctly:
    * `VM_PUBLIC_IP`: Public IP of the server hosting the blockchain network.
    * `VM_USERNAME`: Username used to login to the server hosting the blockchain network.
    * `VM_PASSWORD`: Password used to login to the server hosting the blockchain network.
    * `REDIS_HOST`: URL of the Redis cache.
    * `REDIS_PASSWORD`: Password of the Redis cache.

    More information about setting environment variables can be found [here](https://linuxize.com/post/how-to-set-and-list-environment-variables-in-linux/)

3. Build the image
    ```
    docker build -t ${imageName} -f Dockerfile ${pathToDockerfile} 
    ```


4. Run the Docker image and connects the 8010 ports
    ```
    docker run -it -p 8010:8010 ${imageName}
    ```


<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/Jalmeida1994/GatherChain-Web-Server/issues) for a list of proposed features (and known issues).


<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.



<!-- CONTACT -->
## Contact

João Almeida - [@João Almeida](https://www.linkedin.com/in/jo%C3%A3o-almeida-525476125/) - jcfd.almeida@campus.fct.unl.pt

Project Link: [https://github.com/Jalmeida1994/GatherChain-Web-Server](https://github.com/Jalmeida1994/GatherChain-Web-Server)



<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements
* [FCT-UNL](https://www.fct.unl.pt/)
* [Professor Vasco Amaral](https://docentes.fct.unl.pt/vma/)


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/Jalmeida1994/GatherChain-Web-Server.svg?style=for-the-badge
[contributors-url]: https://github.com/Jalmeida1994/GatherChain-Web-Server/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/Jalmeida1994/GatherChain-Web-Server.svg?style=for-the-badge
[forks-url]: https://github.com/Jalmeida1994/GatherChain-Web-Server/network/members
[stars-shield]: https://img.shields.io/github/stars/Jalmeida1994/GatherChain-Web-Server.svg?style=for-the-badge
[stars-url]: https://github.com/Jalmeida1994/GatherChain-Web-Server/stargazers
[issues-shield]: https://img.shields.io/github/issues/Jalmeida1994/GatherChain-Web-Server.svg?style=for-the-badge
[issues-url]: https://github.com/Jalmeida1994/GatherChain-Web-Server/issues
[license-shield]: https://img.shields.io/github/license/Jalmeida1994/GatherChain-Web-Server.svg?style=for-the-badge
[license-url]: https://github.com/Jalmeida1994/GatherChain-Web-Server/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/jo%C3%A3o-almeida-525476125/
