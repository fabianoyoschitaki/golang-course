$('#new-post-form').on('submit', createPost);

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
        alert('Fail to create post!');
    })
}

// we have data-post-id in the jumbotron div. we get the post id from there
function likePost(event){
    const clickedElement = $(event.target); // getting the <i> that was clicked (posts.html -> "likes" template)
    const postId = clickedElement.closest('div').data('post-id'); // getting closest parent <div> element. we don't use data-post-id, but rather post-id
    console.log("Liking post ID: " + postId);

    // to avoid concurrency in the backend because of multiple clicks, we can disable it
    clickedElement.prop('disabled', true);

    $.ajax({
        url: `/posts/${postId}/likes`,
        method: "POST"
    }).done(function(){
        // alert(`Post ${postId} received a like!`);
        
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
        alert(`Failed to like post ${postId}!`);
    }).always(function(){
        // when ajax request finishes, we enable it again
        clickedElement.prop('disabled', false);
    })
}

function unlikePost(event){
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
        alert(`Failed to unlike post ${postId}!`);
    }).always(function(){
        clickedElement.prop('disabled', false);
    })
}