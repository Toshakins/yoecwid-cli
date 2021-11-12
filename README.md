# yoecwid-cli

Tool to create a store for live preview usage.

Outputs the id of a created store and owner's email.

```

Usage: yoecwid-cli subcommand options

Available subcommands below

store
  -apikey string
        API key for YOLA_B2C_PROFESSIONAL_TIERED plan (Required)
  -password string
        Account password (Required)
  -template-id string
        Template ID, like 'bookstore' (Required)            
  Example: yoecwid-cli store --apikey key --pasword pass --template-id template0
```