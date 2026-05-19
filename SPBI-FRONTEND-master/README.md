# SigapSultan

This project was generated with [Angular CLI](https://github.com/angular/angular-cli) version 13.2.5.

## Tech Stack

**Client:** Angular v13, TailwindCSS v3

**Server:** NodeJS v18.10.0

## Project Structure

```bash
├── dist
├── node_modules
├── src
│   ├── app
│   │   ├── guards
│   │   ├── interceptors
│   │   ├── pages
│   │   ├── services
│   │   └── shared
│   ├── assets
│   ├── environments
│   └── manifest.webmanifest
├── README.md
├── package.json
├── ngsw-config.json
├── tailwind.config.json
└── .gitignore
```

| Folder/File |  Description  |
|:-----:|:--------|
| `/dist`   | build project result |
| `/node_modules`   |  Installed packages  |
| `app/guards`   | Guards for routing |
| `app/interceptors`   | Interceptor of HTTP request |
| `app/pages`   | Consist of modules/pages for webpage views |
| `app/services`   | Consist of classes that are used for HTTP call or shared function |
| `/assets`   | Images, SCSS and other assets for project |
| `/environments`   | List of environment configuration that are used in project |
| `manifest.webmanifest`   | Configuration of PWA |
| `ngsw-config.json`   | Service Worker configuration |
| `tailwind.config.json`   | Tailwind configuration |

## Development server

Clone the project

```bash
  git clone https://marcopolooo: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@github.com/SENTECH-PANGAN-BI/SPBI-FRONTEND
```

Go to the project directory

```bash
  cd SPBI-FRONTEND
```

Install dependencies

```bash
  npm install
```

Start the server

```bash
  npm run start
```

Navigate to `http://localhost:4200/`. The app will automatically reload if you change any of the source files.

## Progressive Web Apps

A PWA only runs on https and localhost environment. The Angular CLI is limited because the service worker doesn’t work with the ng serve command. To see PWA in action, follow these steps:

1. Install `http-server`
    ```bash
    npm i -g http-server
    ```
2. Run command 
    ```bash
    npm run start-pwa
    ```
3. Open project via URL which http-server given

## Code scaffolding

Run `ng generate component component-name` to generate a new component. You can also use `ng generate directive|pipe|service|class|guard|interface|enum|module`.

## Build

Run `npm run build:dev` to build the project. The build artifacts will be stored in the `dist/` directory. Use the `npm run build:prod` flag for a production build.

## Running unit tests

Run `ng test` to execute the unit tests via [Karma](https://karma-runner.github.io).

## Running end-to-end tests

Run `ng e2e` to execute the end-to-end tests via a platform of your choice. To use this command, you need to first add a package that implements end-to-end testing capabilities.

## Further help

To get more help on the Angular CLI use `ng help` or go check out the [Angular CLI Overview and Command Reference](https://angular.io/cli) page.
