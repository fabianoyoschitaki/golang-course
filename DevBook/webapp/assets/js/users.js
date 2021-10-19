// This is for users.html page.
$("#follow").on('click', follow);
$("#unfollow").on('click', unfollow);

// This is from edit-profile.html page
$("#edit-profile-form").on('submit', editProfile);

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
    })
}