# Widgets

![Widgets](../../assets/ComponentFill.svg)

## Out of the box support

- **Text**: A simple text widget that displays a string.
- **Image**: An image widget that displays an image.
- **RSS**: An RSS widget that displays an RSS feed.
- **Weather**: A weather widget that displays the current weather conditions.
- **Clock**: A clock widget that displays the current time.
- **Calendar**: Calendar with import support from iCal (more formats planned)

## Planned

- **Todo**: Todo widget that displays a list of tasks to complete.
- **Shopping list**: Shopping list widget that displays a list of items to buy.
- **Spending statistics**: Spending statistics widget that displays a summary of spending.
- [**Trumf**](###trumf-integration): Integrate and parse data from Trumf.

---

## Details

### Trumf integration

#### Saldo

Saldo widget that displays the current balance

```http
GET https://platform-rest-prod.ngdata.no/trumf/husstand/saldo
```

```json
{
  "trumfSaldo":331.80,
  "totaltAkkumulertTrumf":4737.31,
  "husstandId":"nnnnnnnnn",
  "sistOppdatert":"2025-06-04T10:01:59.000+02:00"
}
```
