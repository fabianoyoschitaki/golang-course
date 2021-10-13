insert into users(name, nick, email, password) values 
    ("John", "john", "john@devbook.com", "$2a$10$aC.C05KibHbgqvb7CCgDJeCaVfis8RJ/PvBDQfTmCKqbCZ0WEk5mW"),
    ("Jane", "jane", "jane@devbook.com", "$2a$10$aC.C05KibHbgqvb7CCgDJeCaVfis8RJ/PvBDQfTmCKqbCZ0WEk5mW"),
    ("Paul", "paul", "paul@devbook.com", "$2a$10$aC.C05KibHbgqvb7CCgDJeCaVfis8RJ/PvBDQfTmCKqbCZ0WEk5mW"),
    ("Peter", "peter", "peter@devbook.com", "$2a$10$aC.C05KibHbgqvb7CCgDJeCaVfis8RJ/PvBDQfTmCKqbCZ0WEk5mW");

insert into user_followers(user_id, follower_user_id) values 
    (1, 2),
    (1, 3),
    (1, 4),
    (2, 1),
    (2, 3),
    (3, 4);

insert into posts(title, content, author_id) values
("Post from John", "This is a post from John!", 1),
("Post from Jane", "This is a post from Jane!", 2),
("Post from Paul", "This is a post from Paul!", 3),
("Post from Peter", "This is a post from Peter!", 4);