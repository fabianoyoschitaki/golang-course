// This is for users.html page.

$("#follow").on('click', follow);
$("#unfollow").on('click', unfollow);

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