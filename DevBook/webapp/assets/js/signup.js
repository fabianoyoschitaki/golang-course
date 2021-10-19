$('#form-signup').on('submit', createUser);

function createUser(event){
    event.preventDefault(); // avoid refreshing page when posting the form
    console.log("Creating user");

    if ($('#password').val() !== $('#confirm-password').val()){
        Swal.fire('Oops...', 'Passwords don\'t match', 'error');
        // alert("Passwords don't match");
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
        // alert("User was successfully created: " + data);
        Swal.fire('Success!', 'User successfully created!', 'success')
            .then(function(){ // after user is created, we want to login the user and redirect him to /home
                $.ajax({
                    url: "/login",
                    method: "POST",
                    data: {
                        email: $('#email').val(),
                        password: $('#password').val()
                    }
                }).done(function(){
                    // we redirect to /home because at this point we'll have the authentication cookies set!
                    window.location = "/home";
                }).fail(function(){
                    Swal.fire('Oops...', 'Error when authentication user!', 'error');
                });
            });

    }).fail(function(error){ // http 40X, 50X
        console.log(error);
        // alert("Fail to create user :(");
        Swal.fire('Oops...', 'Error when creating new user!', 'error');
    })
}