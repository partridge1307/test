CREATE TABLE users (
  	id int PRIMARY KEY,
  	username VARCHAR(255) NOT NULL,
  	password VARCHAR NOT NULL,
  	fullname VARCHAR NOT NULL,
  	avatar VARCHAR,
  	birthday TIMESTAMP,
  	created_time TIMESTAMP NOT NULL
);
  CREATE TYPE status_type as ENUM ('pending', 'accepted', 'rejected');
  
  CREATE TABLE friends (
  	id int PRIMARY KEY,
  	sender_id int NOT NULL,
  	receiver_id int NOT NULL,
  	status STATUS_TYPE,
  	created_time TIMESTAMP,
    
    CONSTRAINT fk_friend_user_1 FOREIGN KEY (sender_id) REFERENCES users(id),
    CONSTRAINT fk_friend_user_2 FOREIGN KEY (receiver_id) REFERENCES users(id)
  );
  
  CREATE TYPE message_type as ENUM ('text', 'image', 'video', 'file');
  CREATE TYPE message_status as ENUM ('sent', 'pending_read', 'read');
  
  CREATE TABLE messages (
  	id int PRIMARY KEY,
    sender_id int NOT NULL,
    receiver_id int NOT NULL,
    type MESSAGE_TYPE NOT NULL,
    status MESSAGE_STATUS,
    content VARCHAR,
    created_time TIMESTAMP,
    
    CONSTRAINT fk_message_user_1 FOREIGN KEY (sender_id) REFERENCES users(id),
    CONSTRAINT fk_message_user_2 FOREIGN KEY (receiver_id) REFERENCES users(id)
  );
  
  
  INSERT INTO users (id, username, password, fullname, created_time) VALUES
  (1, 'partridge', '12345678', 'Nguyen Anh Duc', '2024-08-27 14:10'),
  (2, 'partridge2', '12345678', 'Nguyen Anh Duc', '2024-08-27 14:11'),
  (3, 'partridge3', '12345678', 'Nguyen Anh Duc', '2024-08-27 14:12'),
  (4, 'partridge4', '12345678', 'Nguyen Anh Duc', '2024-08-27 14:13');
  
  INSERT INTO friends (id, sender_id, receiver_id, status, created_time) VALUES
  (1, 1, 2, 'accepted', '2024-08-27 14:10'),
  (2, 1, 3, 'pending', '2024-08-27 14:11'),
  (3, 1, 4, 'rejected', '2024-08-27 14:12'),
  (4, 2, 3, 'accepted', '2024-08-27 14:13');
  
  INSERT INTO messages (id, sender_id, receiver_id, type) VALUES
  (1, 1, 2, 'text'),
  (2, 1, 3, 'image'),
  (3, 1, 4, 'video'),
  (4, 2, 3, 'file');

  SELECT id, username, fullname, avatar FROM users WHERE id = 2 OR id = 3;
  SELECT u.id, u.username, u.fullname, u.avatar FROM users as u JOIN friends as f ON (u.id = f.sender_id OR u.id = f.receiver_id) WHERE (f.sender_id = 2 OR f.receiver_id = 2) AND u.id != 2 AND f.status = 'accepted';
  SELECT * FROM friends as f JOIN messages as m ON f.sender_id = m.sender_id;
