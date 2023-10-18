DROP TABLE IF EXISTS posts;
CREATE TABLE post (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title TEXT NOT NULL,
  'text' TEXT NOT NULL
);


DROP TABLE IF EXISTS tickets;
CREATE TABLE ticket (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  start_date DATE DEFAULT CURRENT_TIMESTAMP, -- right now bitch
  finish_date DATE,
  hourly_rate REAL,
  total_hours REAL,
  client_id INTEGER,
  title TEXT NOT NULL,
  description TEXT NOT NULL, 
  FOREIGN KEY (client_id) REFERENCES client(id)
);

-- load ticket_clock_work by ticket id and compute total time
DROP TABLE IF EXISTS ticket_clock_work;
CREATE TABLE ticket_clock_work (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  ticket_id INTEGER,
  clock_in DATE,
  clock_out DATE,
  hours TIME,
  note TEXT, 
  FOREIGN KEY (ticket_id) REFERENCES ticket(id)
);


DROP TABLE IF EXISTS notes;
CREATE TABLE notes (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  ticket_id INTEGER,
  title TEXT,
  note TEXT,
  FOREIGN KEY (ticket_id) REFERENCES tickets(id)
);


DROP TABLE IF EXISTS clients;
CREATE TABLE client (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  contact_id INTEGER,
  address_id INTEGER
  FOREIGN KEY (contact_id) REFERENCES contact(id),
  FOREIGN KEY (address_id) REFERENCES address(id)
);


DROP TABLE IF EXISTS contacts;
CREATE TABLE contact (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  first_name TEXT NOT NULL,
  last_name TEXT,
  phone TEXT NOT NULL,
  email TEXT,
);


DROP TABLE IF EXISTS addresses;
CREATE TABLE address (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  street_1 TEXT NOT NULL,
  street_2 TEXT, 
  city TEXT NOT NULL,
  state TEXT NOT NULL,
  zip TEXT NOT NULL,
);


DROP TABLE IF EXISTS users;
CREATE TABLE user (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  role_id INTEGER NOT NULL, 
  username TEXT NOT NULL,
  pass_hash TEXT NOT NULL,
  last_active 
  FOREIGN KEY (role_id) REFERENCES user_roles(id)
);


DROP TABLE IF EXISTS user_roles;
CREATE TABLE user_roles (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  role TEXT NOT NULL,
  description TEXT
);


-- to create or to not create a table to monitor requests
