# Capturing high level notes and thoughts

* model.OrderComponent.ComponentID must be values from predefined enums; This way we can check if the order components received are actually products we support and will help in queries as well
* All methods in order_components_repository can be moved under order_repository; I have kept it in separate file for ease of reading and updating
* Move all string constants like error messages to constants file
* Function names and structs should not have reference to AWS or SES - these are just tools we are using right now and susceptible to change
* A good practice for DB migration would be to add updatedAt field to orders table after EKS deplopyment

# Articles for reference
* https://brainhub.eu/library/go-web-app-start
* https://levelup.gitconnected.com/a-practical-approach-to-structuring-go-applications-7f77d7f9c189
* https://tutorialedge.net/golang/go-project-structure-best-practices/
* https://articles.wesionary.team/a-clean-architecture-for-web-application-in-go-lang-4b802dd130bb
* https://techinscribed.com/different-approaches-to-pass-database-connection-into-controllers-in-golang/

# Code templates already available with SOLID desing principles
* https://github.com/irahardianto/service-pattern-go
* https://github.com/qiangxue/go-rest-api ---- This is the best template I have found for our requirements which has logging, authentication middleware and unit testing