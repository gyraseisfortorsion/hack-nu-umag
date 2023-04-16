ALTER TABLE umag_hacknu.sale
CHANGE COLUMN sale_time timeSt DATETIME;

ALTER TABLE umag_hacknu.supply
CHANGE COLUMN supply_time timeSt DATETIME;

ALTER TABLE umag_hacknu.sale ADD COLUMN type INT DEFAULT 0;
ALTER TABLE umag_hacknu.supply ADD COLUMN type INT DEFAULT 1;

CREATE TABLE umag_hacknu.sale_new AS
SELECT id, barcode, quantity, price, time, type
FROM umag_hacknu.sale
UNION ALL
SELECT id, barcode, quantity, price, time, type
FROM umag_hacknu.supply;

ALTER TABLE umag_hacknu.sale_new ADD COLUMN cogs int(11) DEFAULT 0;
ALTER TABLE umag_hacknu.sale_new ADD COLUMN revenue int(11) DEFAULT 0;
ALTER TABLE umag_hacknu.sale_new ADD COLUMN profit int(11) DEFAULT 0;
ALTER TABLE umag_hacknu.sale_new ADD COLUMN cumulative_profit int(11) DEFAULT 0;