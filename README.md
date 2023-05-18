# GoTestAPI

# 1. API
The repository contains a basic CRUD API written in GoLang. This project was aimed to perform blackbox testing on this API. 

The following CRUD operations are been added:
1. Create - createRoll.go
2. Read - getRoll.go and getRolls.go
3. Update - updateRoll.go
4. Delete - deleteRoll.go

# createRoll
Creates a new item in the database.

# getRoll
Fetch information for a particular item present in the database.

# getRolls
Fetch information for all the itmes present in the database.

# updateRoll
Update the information for a particular item present in the database.

# deleteRoll
Delete the information and item itself from the database.

# 2. Client API’s
I have used the API to perform CRUD operations on "CLIENT" dataset and hence renamed them as follows:

# 2.1 Get Clients API
![image](https://github.com/heygauri/GoTestAPI/assets/64316529/4e08945a-3947-4863-9e21-4b3571621e8b)


# 2.2 Get Client API
![image](https://github.com/heygauri/GoTestAPI/assets/64316529/b247568c-0fb7-4fd1-958e-3e219cefdf81)


# 2.3 Create Client API
![image](https://github.com/heygauri/GoTestAPI/assets/64316529/3d4d686c-fcde-4b81-859b-265efba929e4)


# 2.4 Update Client API
![image](https://github.com/heygauri/GoTestAPI/assets/64316529/3590b9aa-d21b-4ca5-85a0-e8dc9c72b200)


# 2.5 Delete Client API
![image](https://github.com/heygauri/GoTestAPI/assets/64316529/960e7f80-02d5-4c67-8d14-093f2c57a36e)



# 3. Class Diagram for API’s

![image](https://github.com/heygauri/GoTestAPI/assets/64316529/502ebd24-118d-4817-8a64-9e84ca0a44c0)


# 4. DATABASE
NOTE: For this project I have used the local database itself because the goal is to perform the BlackBox Testing on this API. So the main focus is on testing here. 

The variable "clients" is a slice and is initialized to store the database in main.go file.

# 5. TESTING
For testing I have built a Static Design Documentation for Test Cases "Static_Design.pdf".

Following the Static design documenation various test cases are build. Some of them are stored in "Test Cases for API Testing - Sheet1.pdf". In the sheet the red color denotes the failed test case and the green one is the updated version of the failed test case.

# 5.1 testing tool
The testing tool used for API testing is POSTMAN. It is used for both manual and automated API testing. 

# 5.2 Manual testing
To perform MANUAL testing refer to "Test Cases for API Testing - Sheet1.pdf" for pre-build test cases and "Static_Design.pdf" to build new test cases.

# 5.3 Automated Testing
To perform AUTOMATED testing use the following scripts in postman.
const clientBaseUrl = pm.variables.get("clientBaseUrl");

# 5.3.1 GetClients
![image](https://github.com/heygauri/GoTestAPI/assets/64316529/2f04cdb8-f36b-41c0-a803-a4780d64c96b)

The test script for GetClients is "getClientsTestScript.js".

# 5.3.2 GetClient
![image](https://github.com/heygauri/GoTestAPI/assets/64316529/1068b992-46cc-4619-8a35-b2746bbe0374)

The test script for GetClient is "getClientTestScript.js".

# 5.3.3 CreateClient
![image](https://github.com/heygauri/GoTestAPI/assets/64316529/437d8fe3-7100-41c2-8318-7ce5f487960f)

The test script for CreateClient is "createClientTestScript.js".

# 5.3.4 UpdateClient
![image](https://github.com/heygauri/GoTestAPI/assets/64316529/69f3f85f-7c3b-410f-827e-6274d4f6f5dc)

The test script for UpdateClient is "updateClientTestScript.js".

# 5.3.5 DeleteClient
![image](https://github.com/heygauri/GoTestAPI/assets/64316529/574b838e-e874-4f48-a534-4fa867463eb6)

The test script for DeleteClient is "deleteClientTestScript.js".











