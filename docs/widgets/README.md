# Widgets

![Widgets](../../assets/ComponentFill.svg)


- :lightning:: Dynamic widget


Widget type | Type | Description |
|-----------|------|-------------|
| Static | Text | A simple text widget that displays a string. |
| Static | Image | An image widget that displays an image. |
| Static | RSS | An RSS widget that displays an RSS feed. |
| Static | Weather | A weather widget that displays the current weather conditions. |
| :lightning: | Lighting radar | A lighting radar widget that displays the current lighting conditions. |
| Static | Clock | A clock widget that displays the current time. |
| Static | Calendar | Calendar with import support from iCal (more formats planned) |

## Out of the box support

- **Text**: A simple text widget that displays a string.
- **Image**: An image widget that displays an image.
- **RSS**: An RSS widget that displays an RSS feed.
- **Weather**: A weather widget that displays the current weather conditions.
  - :lightning: Lighting radar**
- **Clock**: A clock widget that displays the current time.
- **Calendar**: Calendar with import support from iCal (more formats planned)

## Planned

- **Todo**: Todo widget that displays a list of tasks to complete.
- **Shopping list**: Shopping list widget that displays a list of items to buy.
- **Spending statistics**: Spending statistics widget that displays a summary of spending.
- [Trumf](#Trumf-integration): Integrate and parse data from Trumf.

---


## Calendar

Automatically fetch calendars from pre-defined sources such as [Spond](https://www.spond.com/)

- [Spond API](https://api.spond.com/core/v1/)
- [Olen/Spond](https://github.com/Olen/Spond): Relevant repo

Pseudo code for future reference

```go
import "pkg/wasmdash/middleware/secrets"

// Data type
type SpondAPI struct {
	Groups []Group
	Messages []Message
	Events []Event

	Auth secrets.PasswordAuth
}

// Group type
type Group struct {
	ID   string
	Name string
	Calendar Calendar
}
...
}
```

## Spending statistics

Fetch data from Trumf and Rema.

Trumf data is easy to gather as they provide an API, but for Rema, it's not quite so straight forward.
- Needs research for automating this.


Pseudo code for future reference

```go
import "pkg/wasmdash/middleware/secrets"

type TrumfData struct {}
type RemaData struct {}

type TrumfAPI struct {
	Auth secrets.PasswordAuth
}

type RemaAPI struct {
	Auth secrets.PasswordAuth
}
```

### Trumf integration

#### Saldo

Saldo widget that displays the current balance

```http
GET https://platform-rest-prod.ngdata.no/trumf/husstand/saldo
```

```json
{
  "trumfSaldo":1234.56,
  "totaltAkkumulertTrumf":1234567.89,
  "husstandId":"nnnnnnnnn",
  "sistOppdatert":"2025-06-04T10:01:59.000+02:00"
}
```
