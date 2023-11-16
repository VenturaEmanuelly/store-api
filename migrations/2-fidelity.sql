CREATE TABLE fidelity (
    id INT,
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    plan VARCHAR,

    FOREIGN KEY(id) REFERENCES registration (id),
    UNIQUE(id)
)