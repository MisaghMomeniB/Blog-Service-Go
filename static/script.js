// Fetch and display all blog posts
document.getElementById("loadPosts").addEventListener("click", function() {
    fetch("/posts")
        .then(response => response.json())
        .then(data => {
            let postsList = document.getElementById("postsList");
            postsList.innerHTML = "";
            data.forEach(post => {
                let postDiv = document.createElement("div");
                postDiv.classList.add("post");
                postDiv.innerHTML = `<h3>${post.title}</h3><p>${post.content}</p><small>Created at: ${new Date(post.created_at).toLocaleString()}</small><button onclick="deletePost(${post.id})">Delete</button>`;
                postsList.appendChild(postDiv);
            });
        })
        .catch(error => {
            console.error("Error loading posts:", error);
        });
});

// Handle post creation
document.getElementById("createPostForm").addEventListener("submit", function(event) {
    event.preventDefault();

    const title = document.getElementById("title").value;
    const content = document.getElementById("content").value;

    const post = { title, content };

    fetch("/posts/create", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(post)
    })
    .then(response => response.json())
    .then(data => {
        alert("Post created successfully!");
        document.getElementById("title").value = '';
        document.getElementById("content").value = '';
    })
    .catch(error => {
        console.error("Error creating post:", error);
    });
});

// Handle post deletion
function deletePost(postId) {
    fetch(`/posts/delete?id=${postId}`, {
        method: "DELETE",
    })
    .then(response => {
        if (response.ok) {
            alert("Post deleted!");
            document.getElementById("loadPosts").click();
        } else {
            alert("Failed to delete post.");
        }
    })
    .catch(error => {
        console.error("Error deleting post:", error);
    });
}
