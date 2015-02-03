(function () {
    var baseAuthHash = "";
    $.ajaxSetup({
        beforeSend: function(xhr) {
            xhr.setRequestHeader ("Authorization", "Basic " + baseAuthHash);
        }
    });
    
    function setBaseAuth(user, password) {
        var tok = user + ':' + password;
        baseAuthHash = Base64.encode(tok);
    }
    
    $(function() {
        $("#usersForm").on("submit", onUsersFormSubmit)
        $("#loginForm").on("submit", onLoginFormSubmit)
    });
    
    function onLoginFormSubmit(event) {
        event.preventDefault();
        
        setBaseAuth($("#name").val(), $("#password").val());
        
        $("#loginError").text("");
        
        var data = $(this).serialize()
        
        $.get("/accounts/checkAuth")
            .done(function(data) {
                if(data.error){
                    $("#loginError").text(data.error).show();
                } else {
                    $("#login").hide();
                    showUsers();
                }
            });
            
        return false
    }
    
    function showUsers() {
        $("#users").show();
        loadUsers();    
    }
    
    function loadUsers() {
        $("#usersError").text("");
        
        $.get("/users")
            .done(function(data) {
                if(data.error){
                    $("#usersError").text(data.error).show();
                } else {
                    $("#usersTable tbody").empty();
                    
                    var users = data.data;
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
        $("<td>").text(user.username).appendTo(row);
        $("<td>").text(user.firstname).appendTo(row);
        $("<td>").text(user.lastname).appendTo(row);
                
        $("#usersTable tbody").append(row);
    }
    
    function onUsersFormSubmit(event) {
        event.preventDefault();
        
        $("#usersError").text("");
        
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