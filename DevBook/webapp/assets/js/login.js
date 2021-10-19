$('#form-login').on('submit', attemptLogin);

function attemptLogin(event){
    event.preventDefault(); // avoid refreshing page when posting the form
    console.log("Attempting login");

    $.ajax({
        url: "/login", // this is not the API route, but webapp route which will call the API backend route
        method: "POST",
        data: {
            email: $('#email').val(),
            password: $('#password').val()
        }
    }).done(function(data){ // http 20X
        console.log("Login was successful!");
        // let's redirect the user to his feed
        window.location = "/home";
    }).fail(function(error){ // http 40X, 50X
        console.log(error);
        // alert("Fail to login :(");
        Swal.fire('Oops...', 'Email or password are incorrect!', 'error');
    })
}