// This is for users.html page.
$("#follow").on('click', follow);
$("#unfollow").on('click', unfollow);

// This is from edit-profile.html page
$("#edit-profile-form").on('submit', editProfile);

// This is from change-password.html
$("#change-password-form").on('submit', updatePassword);

// This is from profile.html page (to delete user)
$("#delete-user").on('click', deleteUser);

function unfollow(){
    console.log("Unfollow");
    const userToUnfollowId = $(this).data('user-id');
    $(this).prop('disabled', true); // to avoid multiple clicks
    $.ajax({
        url: `/users/${userToUnfollowId}/unfollow-user`,
        method: "POST"
    }).done(function(){
        // to avoid lots of refreshes (e.g. changing button to follow, updating followers etc, let's just refresh the page)
        window.location = `/users/${userToUnfollowId}`;
    }).fail(function(){
        Swal.fire('Oops...', 'Error to unfollow user!', 'error');
        $("#unfollow").prop('disabled', false);
    })
}

function follow(){
    console.log("Follow");
    const userToFollowId = $(this).data('user-id');
    $(this).prop('disabled', true); // to avoid multiple clicks
    $.ajax({
        url: `/users/${userToFollowId}/follow-user`,
        method: "POST"
    }).done(function(){
        // to avoid lots of refreshes (e.g. changing button to follow, updating followers etc, let's just refresh the page)
        window.location = `/users/${userToFollowId}`;
    }).fail(function(){
        Swal.fire('Oops...', 'Error to follow user!', 'error');
        $("#follow").prop('disabled', false);
    })
}

// we need event to prevent default
function editProfile(event){
    event.preventDefault();
    console.log("Editing profile...");
    $.ajax({
        url: '/edit-profile',
        method: 'PUT',
        data: {
            newName: $('#new-name').val(),
            newEmail: $('#new-email').val(),
            newNick: $('#new-nick').val(),
        }
    }).done(function(){
        Swal.fire('Success!', 'User data was successfully updated!', 'success')
            .then(function(){
                window.location = '/profile';
            });
    }).fail(function(){
        Swal.fire('Oops...', 'Error to update profile :(', 'error');
    });
}

// this is for change-password.html
function updatePassword(event){
    console.log("Updating password")
    event.preventDefault();

    // validate if user provided current password
    const currentPassword = $("#current-password").val();
    if (!currentPassword){
        Swal.fire('Oops...', 'You need to inform your current password!', 'warning');
        return;
    }

    // validate if new password is != confirmation password
    const newPassword = $("#new-password").val();
    const confirmationPassword = $("#confirm-password").val();

    if (newPassword !== confirmationPassword){
        console.warn("New password is [" + newPassword + "] but confirmation password is [" + confirmationPassword + "]")
        Swal.fire('Oops...', 'Passwords don\'t match!', 'warning');
        return;
    }

    // if they're equals, let's call FE route
    $.ajax({
        url: '/change-password',
        method: "POST",
        data: {
            "currentPassword": currentPassword,
            "newPassword": newPassword
        }
    }).done(function(){
        Swal.fire('Success!', 'Your password was successfully changed!', 'success')
            .then(function(){
                window.location = '/profile';
            });
    }).fail(function(){
        Swal.fire('Oops...', 'Error to change password :(', 'error');
    });
}

function deleteUser(){
    // before actually deleting, we ask the user if he's sure about that 
    Swal.fire({
        title: 'Attention!',
        text: 'Are you sure you want to delete your account? This change is permanent.',
        showCancelButton: true,
        cancelButtonText: 'Cancel',
        icon: 'warning'
    }).then(function(confirmation){ // runs when swal is closed
        // if user cancelled
        if (!confirmation.value){
            return;
        }
        
        $.ajax({
            url: '/delete-user',
            method: "DELETE"
        }).done(function(){
            // we show the user his account was deleted and then we log him out
            Swal.fire('success', 'Your account was succesfully deleted!', 'success')
                .then(function(){
                    window.location = '/logout';
                });

        }).fail(function(error){
            Swal.fire('Oops...', 'Failed to delete your account!', 'error');
        });
    });
}