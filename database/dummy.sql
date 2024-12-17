-- Aktifkan ekstensi uuid-ossp
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Tabel Users
INSERT INTO users (id, name, email, password, role)
VALUES
  (uuid_generate_v4(), 'Budi Santoso', 'budi.santoso@email.com', 'password123', 'customer'),
  (uuid_generate_v4(), 'Siti Aminah', 'siti.aminah@email.com', 'password123', 'customer'),
  (uuid_generate_v4(), 'Agus Prasetyo', 'agus.prasetyo@email.com', 'password123', 'admin'),
  (uuid_generate_v4(), 'Dewi Lestari', 'dewi.lestari@email.com', 'password123', 'customer');

-- Tabel Addresses
INSERT INTO addresses (id, user_id, address, city, postal_code)
VALUES
  (uuid_generate_v4(), (SELECT id FROM users WHERE name = 'Budi Santoso'), 'Jl. Merdeka No. 10', 'Jakarta', '10110'),
  (uuid_generate_v4(), (SELECT id FROM users WHERE name = 'Siti Aminah'), 'Jl. Asia Afrika No. 20', 'Bandung', '40123'),
  (uuid_generate_v4(), (SELECT id FROM users WHERE name = 'Agus Prasetyo'), 'Jl. Diponegoro No. 30', 'Surabaya', '60241'),
  (uuid_generate_v4(), (SELECT id FROM users WHERE name = 'Dewi Lestari'), 'Jl. Ahmad Yani No. 40', 'Yogyakarta', '55122');

-- Tabel Shipments
INSERT INTO shipments (id, customer_id, origin_address_id, destination_address_id, status_id, service_type, created_at, status)
VALUES
  (uuid_generate_v4(), 
   (SELECT id FROM users WHERE name = 'Budi Santoso'),
   (SELECT id FROM addresses WHERE address = 'Jl. Merdeka No. 10'),
   (SELECT id FROM addresses WHERE address = 'Jl. Asia Afrika No. 20'),
   uuid_generate_v4(), 'Reguler', NOW(), 'Dikirim'),

  (uuid_generate_v4(), 
   (SELECT id FROM users WHERE name = 'Siti Aminah'),
   (SELECT id FROM addresses WHERE address = 'Jl. Asia Afrika No. 20'),
   (SELECT id FROM addresses WHERE address = 'Jl. Ahmad Yani No. 40'),
   uuid_generate_v4(), 'Express', NOW(), 'Diterima');

-- Tabel Shipment Details
INSERT INTO shipment_details (id, shipment_id, weight, description)
VALUES
  (uuid_generate_v4(), 
   (SELECT id FROM shipments WHERE service_type = 'Reguler'),
   3.5, 'Peralatan Dapur'),
  
  (uuid_generate_v4(), 
   (SELECT id FROM shipments WHERE service_type = 'Reguler'),
   1.2, 'Buku Pelajaran'),
  
  (uuid_generate_v4(), 
   (SELECT id FROM shipments WHERE service_type = 'Express'),
   2.0, 'Elektronik Kecil');

-- Tabel Pricing
INSERT INTO pricings (id, origin, destination, weight, service_type, price)
VALUES
  (uuid_generate_v4(), 'Jakarta', 'Bandung', 4.7, 'Reguler', 75000.00),
  (uuid_generate_v4(), 'Bandung', 'Yogyakarta', 2.0, 'Express', 120000.00),
  (uuid_generate_v4(), 'Surabaya', 'Jakarta', 5.0, 'Reguler', 95000.00);

-- Tabel Payments
INSERT INTO payments (id, shipment_id, amount, payment_status, paid_at)
VALUES
  (uuid_generate_v4(), 
   (SELECT id FROM shipments WHERE service_type = 'Reguler'), 
   75000.00, 'Lunas', NOW()),

  (uuid_generate_v4(), 
   (SELECT id FROM shipments WHERE service_type = 'Express'), 
   120000.00, 'Lunas', NOW());
