CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  email VARCHAR(150) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  role VARCHAR(20) DEFAULT 'user',
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE category (
  id SERIAL PRIMARY KEY,
  category_name VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE product (
  id SERIAL PRIMARY KEY,
  product_name VARCHAR(150) NOT NULL,
  description TEXT,
  created_at TIMESTAMP DEFAULT NOW(),
  price FLOAT NOT NULL,
  quantity INT DEFAULT 0,
  in_stock BOOLEAN DEFAULT true,
  category_id INT
);

CREATE TABLE cart_items (
  id SERIAL PRIMARY KEY,
  user_id INT,
  product_id INT,
  qty INT CHECK (qty > 0),
  price FLOAT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE orders (
  id SERIAL PRIMARY KEY,
  user_id INT,
  created_at TIMESTAMP DEFAULT NOW(),
  status VARCHAR(20) DEFAULT 'placed'
);

CREATE TABLE order_item (
  id SERIAL PRIMARY KEY,
  order_id INT,
  product_id INT,
  quantity INT CHECK (quantity > 0),
  price FLOAT NOT NULL
);

-- Relationships
ALTER TABLE product ADD CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES category (id);
ALTER TABLE cart_items ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE cart_items ADD CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES product (id);
ALTER TABLE orders ADD CONSTRAINT fk_order_user FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE order_item ADD CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES orders (id);
ALTER TABLE order_item ADD CONSTRAINT fk_order_product FOREIGN KEY (product_id) REFERENCES product (id);
