
-- Insert User
    INSERT INTO
    `gcp_project_id.money_minder.t_users` (ID,
        Name,
        Email,
        Password,
        CreatedAt)
    VALUES
    (
        1, 'Sounish Nath', 'flock.sinasini@gmail.com', 'xxxxx', CURRENT_TIMESTAMP()
    );

-- Insert Category
    INSERT INTO `gcp_project_id.money_minder.t_categories`
    (ID, Category, IsActive)
    VALUES
    (8, 'xx', true)

-- Add Transaction
    INSERT INTO `gcp_project_id.money_minder.t_transactions`
    (ID,Name,TransactionType,UserID,CategoryID,Amount,CreatedAt)
    VALUES
    (2, 'Mumbai House', 'Expense', 1, 10, 14000, current_timestamp());