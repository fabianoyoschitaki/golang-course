$('#form-signup').on('submit', createUser);

function createUser(event){
    event.preventDefault(); // avoid refreshing page when posting the form
    console.log("Creating user");

    if ($('#password').val() !== $('#confirm-password').val()){
        alert("Passwords don't match");
        return;
    }

    $.ajax({
        url: "/users", // this is not the API route, but webapp route which will call the API backend route
        method: "POST",
        data: {
            name: $('#name').val(),
            nick: $('#nick').val(),
            email: $('#email').val(),
            password: $('#password').val()
        }
    }).done(function(data){ // http 20X
        console.log("User was successfully created!");
        alert("User was successfully created: " + data);
    }).fail(function(error){ // http 40X, 50X
        console.log(error);
        alert("Fail to create user :(");
    })
}