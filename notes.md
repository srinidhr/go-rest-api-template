# Capturing high level notes and thoughts

* model.OrderComponent.ComponentID must be values from predefined enums; This way we can check if the order components received are actually products we support and will help in queries as well
* All methods in order_components_repository can be moved under order_repository; I have kept it in separate file for ease of reading and updating
* Move all string constants like error messages to constants file
* Function names and structs should not have reference to AWS or SES - these are just tools we are using right now and susceptible to change
* A good practice for DB migration would be to add updatedAt field to orders table after EKS deplopyment