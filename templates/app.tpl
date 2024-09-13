package main

import (
    "{{.Module}}/controllers"
    "{{.Module}}/services"
)

func SetupApp() {
    // Initialize your app here
    controllers.InitControllers()
    services.InitServices()
}
