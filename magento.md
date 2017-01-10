migrate from magento
---

* Table storage engine for 'catalog_product_website' doesn't have this option
```bash
sed -i 's/ROW_FORMAT=FIXED//g' change-me.sql
```
