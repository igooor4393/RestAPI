CREATE TABLE requests (
                          id SERIAL PRIMARY KEY,
                          requestType VARCHAR(255) NOT NULL,
                          input VARCHAR(255) NOT NULL,
                          output VARCHAR(255) NOT NULL
                      );