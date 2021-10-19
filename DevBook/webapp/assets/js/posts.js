// from home.html
$('#new-post-form').on('submit', createPost);

// from update-post.html
$("#btn-update-post").on('click', updatePost);

// from home.html
$(".delete-post-icon").on('click', deletePost);

// we need to make this way for the cases where the element get classes changed during runtime
$(document).on('click', '.like-post', likePost);
$(document).on('click', '.unlike-post', unlikePost);

// calling webapp route to create new post
function createPost(event){
    event.preventDefault();
    $.ajax({
        url: "/posts",
        method: "POST",
        data: {
            newPostTitle: $('#title').val(),
            newPostContent: $('#content').val()
        }
    }).done(function(){
        window.location = "/home"
    }).fail(function(){
        Swal.fire('Oops...', `Failed to create new post!`, 'error');
    })
}

// we have data-post-id in the jumbotron div. we get the post id from there
function likePost(event){
    event.preventDefault();

    const clickedElement = $(event.target); // getting the <i> that was clicked (posts.html -> "likes" template)
    const postId = clickedElement.closest('div').data('post-id'); // getting closest parent <div> element. we don't use data-post-id, but rather post-id
    console.log("Liking post ID: " + postId);

    // to avoid concurrency in the backend because of multiple clicks, we can disable it
    clickedElement.prop('disabled', true);

    $.ajax({
        url: `/posts/${postId}/likes`,
        method: "POST"
    }).done(function(){
        // getting element that has the number of likes
        const likeCounter = clickedElement.next('span');

        // getting the actual number of likes
        const likeCount = parseInt(likeCounter.text());

        // refreshing the number of likes to +1
        likeCounter.text(likeCount + 1);

        // in our case, the unlike feature will be enabled when we like a post by changing its function. 
        // given the scope of the course, it only works in case the user does not refresh the page
        clickedElement.addClass('unlike-post');
        clickedElement.addClass('text-danger'); // make it red
        clickedElement.removeClass('like-post');
    }).fail(function(error){
        Swal.fire('Oops...', `Failed to like post ${postId}!`, 'error');
    }).always(function(){
        // when ajax request finishes, we enable it again
        clickedElement.prop('disabled', false);
    })
}

function unlikePost(event){
    event.preventDefault();

    const clickedElement = $(event.target); // getting the <i> that was clicked (posts.html -> "likes" template)
    const postId = clickedElement.closest('div').data('post-id'); // getting closest parent <div> element. we don't use data-post-id, but rather post-id
    console.log("Unliking post ID: " + postId);

    // to avoid concurrency in the backend because of multiple clicks, we can disable it
    clickedElement.prop('disabled', true);

    $.ajax({
        url: `/posts/${postId}/unlikes`,
        method: "POST"
    }).done(function(){
        const likeCounter = clickedElement.next('span');
        const likeCount = parseInt(likeCounter.text());
        likeCounter.text(likeCount - 1);

        clickedElement.removeClass('text-danger'); // make it black again
        clickedElement.removeClass('unlike-post');

        clickedElement.addClass('like-post');
    }).fail(function(error){
        Swal.fire('Oops...', `Failed to unlike post ${postId}!`, 'error');
    }).always(function(){
        clickedElement.prop('disabled', false);
    })
}

function deletePost(event){
    event.preventDefault();

    const deleteIconElement = $(event.target);
    const postDiv = deleteIconElement.closest('div'); 
    const postId = postDiv.data('post-id');

    // before actually deleting, we ask the user if he's sure about that 
    Swal.fire({
        title: 'Attention!',
        text: `Are you sure you want to delete post ${postId}?`,
        showCancelButton: true,
        cancelButtonText: 'Cancel',
        icon: 'warning'
    }).then(function(confirmation){ // runs when swal is closed
        // if cancelled
        if (!confirmation.value){
            return;
        }
        
        // otherwise he confirmed, let's delete post
        deleteIconElement.prop('disabled', true);
        $.ajax({
            url: `/posts/${postId}`,
            method: "DELETE"
        }).done(function(){
            // we remove the jumbotron, don't need to refresh the page. function is what runs when fadeOut finishes
            postDiv.fadeOut("slow", function(){
                $(this).remove();
            });
        }).fail(function(error){
            Swal.fire('Oops...', `Failed to delete post ${postId}!`, 'error');
        });
    });
}

function updatePost(event) {
    // disable save changes button
    $(this).prop('disabled', true);

    // getting data-post-id value
    const postId = $(this).data('post-id');

    const title = $("#title").val();
    const content = $("#content").val();

    console.log("Updating post " + postId + " with new title: " + title + " and content: " + content);
    $.ajax({
        url: `/posts/${postId}`,
        method: "PUT",
        data: {
            "title": title,
            "content": content
        }
    }).done(function(data){
        // this is SweetAlert 
        Swal.fire(
            'Success!',
            `Post ${postId} succesfully updated!`,
            'success'
        ).then(function(){
            window.location = '/home';
        });
    }).fail(function(error){
        Swal.fire('Oops...', `Failed to update post ${postId}!`, 'error');
        console.error(error);
    }).always(function(){
        // $(this) = this function, so we need to use the button id
        $('#btn-update-post').prop('disabled', false);
    });
}