# Capturing high level notes and thoughts

* `model.OrderComponent.ComponentID` must be values from predefined enums; This way we can check if the order components received are actually products we support and will help in queries as well
* All methods in `order_components_repository` can be moved under order_repository; I have kept it in separate file for ease of reading and updating
* Move all string constants like error messages to constants file
* Configuration variables should be loaded at the start of the server into a singleton OR single source within the Go code rather than references it from OS/environment every time within the service
* Function names and structs should not have reference to AWS or SES - these are just tools we are using right now and susceptible to change
* A good practice for DB migration would be to add updatedAt field to orders table after EKS deplopyment
* Logging should be correctly structured with log levels such as info, error, etc and other information such as timestamp, caller info, etc. We should also introduce the ability to trace logs across services. - [Reference](https://blog.logrocket.com/5-structured-logging-packages-for-go/)
* All integration OR e2e tests should clean up the test data after execution

# Articles for reference
* https://brainhub.eu/library/go-web-app-start
* https://levelup.gitconnected.com/a-practical-approach-to-structuring-go-applications-7f77d7f9c189
* https://tutorialedge.net/golang/go-project-structure-best-practices/
* https://articles.wesionary.team/a-clean-architecture-for-web-application-in-go-lang-4b802dd130bb
* https://techinscribed.com/different-approaches-to-pass-database-connection-into-controllers-in-golang/

# Code templates already available with SOLID desing principles
* https://github.com/irahardianto/service-pattern-go
* https://github.com/qiangxue/go-rest-api ---- This is the best template I have found for our requirements which has logging, authentication middleware and unit testing
