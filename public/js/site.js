(function () {
    $(function() {
        loadUsers();
        
        $("#usersForm").on("submit", onUsersFormSubmit)
    });
    
    function loadUsers() {
        $("#usersError").hide();
        
        $.get("/users")
            .done(function(data) {
                if(data.error){
                    $("#usersError").text(data.error).show();
                } else {
                    $("#usersTable tbody").empty();
                    
                    var users = data;
                    for(var i = 0; i < users.length; i++) {
                        var user = users[i];
                        addUserToTable(user);
                    }
                }
            });
    }
    
    function addUserToTable(user) {
        var row = $("<tr>")
        $("<td>").text(user.id).appendTo(row);
        $("<td>").text(user.firstname).appendTo(row);
        $("<td>").text(user.lastname).appendTo(row);
                
        $("#usersTable tbody").append(row);
    }
    
    function onUsersFormSubmit(event) {
        event.preventDefault();
        
        $("#usersError").hide();
        
        var data = $(this).serialize()
        
        $.post("/users", data)
            .done(function(data) {
                if(data.error){
                    $("#usersError").text(data.error).show();
                } else {
                    loadUsers();
                }
            });
            
        return false
    }
})();