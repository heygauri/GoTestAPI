// Set the base URL for the API
pm.environment.set("baseUrl", "http://localhost:5000")

// Test case 1: Try to delete a client that exists
pm.test("Delete an existing client", function () {
    // Send a DELETE request to the API to delete the client with ID 1
    pm.sendRequest({
        url: pm.environment.get("baseUrl") + "/client/2008",
        method: "DELETE",
        headers: {
            "Content-Type": "application/json"
        }
    }, function (err, res) {
        // Check if the response has a status code of 200
        //pm.expect(res.status).to.equal(OK);
        pm.expect(res).to.have.status(200);
        //pm.expect(res.json().error).to.eql("Client deleted successfully");
    });
});

// Test case 2: Try to delete a client that doesn't exist
pm.test("Delete a non-existing client", function () {
    // Send a DELETE request to the API to delete a client with an invalid ID
    pm.sendRequest({
        url: pm.environment.get("baseUrl") + "/client/1000",
        method: "DELETE",
        headers: {
            "Content-Type": "application/json"
        }
    }, function (err, res) {
        // Check if the response has a status code of 404
        pm.expect(res).to.have.status(404);
    });
});

// Test case 3: Try to delete a client with an invalid ID
pm.test("Delete a client with invalid ID", function () {
    // Send a DELETE request to the API to delete a client with an invalid ID
    pm.sendRequest({
        url: pm.environment.get("baseUrl") + "/client/abc",
        method: "DELETE",
        headers: {
            "Content-Type": "application/json"
        }
    }, function (err, res) {
        // Check if the response has a status code of 404
        pm.expect(res).to.have.status("Not Found");
    });
});

// Test case 4: Try to delete a client without specifying the ID
pm.test("Delete a client without specifying ID", function () {
    // Send a DELETE request to the API without specifying the ID of the client to delete
    pm.sendRequest({
        url: pm.environment.get("baseUrl") + "/client/",
        method: "DELETE",
        headers: {
            "Content-Type": "application/json"
        }
    }, function (err, res) {
        // Check if the response has a status code of 404
        pm.expect(res.status).to.equal("Not Found");
    });
});
