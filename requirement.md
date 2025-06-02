Golang Assignment:
 Import and Manage Data
 Objective:
 Build a Golang program to import data from an Excel file, store it into MySQL,
 and cache the data in Redis. Create a simple CRUD (Create, Read, Update,
 Delete) system to view imported data, edit records, and update the changes to
 both the database and cache.
 Framework:
 gin
 Task:
 Import Excel Data:
 ● Uploadgivenexcel file with API development tool like postman
 ● Parsethedataandstructure it appropriately.
 ● Validate the uploaded Excel file format, ensuring it adheres to specific
 column headers and data types.
 ● Implementasynchronousprocessing for parsing and structuring the
 data to improve scalability.
 Store Data:
 ● ConnecttoMySQLandcreateatabletostore the imported data.
 ● Implementafunctionto insert the parsed data into the MySQL database.
 ● ConnecttoRedisandcachetheimported data.
 ● Rediscacheshouldexpire after 5 minus.
ViewImportedList:
 ● CreateanAPIendpointoracommand-lineinterface to view imported
 data fetched from Redis. If Redis doesn't have the data, retrieve it from
 the table
 ● Displaythe data in a readable format.
 Edit Record:
 ● Allowuserstoedit aspecific record.
 ● UpdatetherecordinboththeMySQLdatabaseandRediscache.
 Handle Errors andValidation:
 ● Implementerrorhandling mechanisms to gracefully handle failures
 during file upload, data parsing, database operations, etc.
 ● Validate the uploaded Excel file to ensure it meets the expected format
 and structure.
 Optimization and Scalability:
 ● Optimizedatabase queries and Redis operations for better performance.
 ● Designtheapplication architecture to scale horizontally to handle
 increased traffic and data volume efficiently.
 Submission:
 ● ProvidetheGolangcodealong with necessary configuration files.
 ● IncludeaREADME.mdwithinstructions if any
● Note:Ensurethat the necessary Golang packages for Excel parsing,
 MySQL, andRedis are used, and all dependencies are clearly
 documented.
 ● Thecodeshouldbewell-organized, follow best practices, and include
 error handling mechanisms