# Documentation

## Theme middleware

In a templ file, you can use the theme middleware to apply a theme to your components. The theme middleware is a function that takes a component and returns a new component with the theme applied.

```templ
// Dashboard-themed button
@button.Button(button.Props{
    Variant: button.VariantDashboard, // Variants are used to define the appearance of a button,
                                      // and can be added and modified in pkg/components/button.templ
    Class:   "w-full",
}) {
    Dashboard Action
}

// Accent card
@card.Card(card.Props{
    Class: "border-dashboard-accent",
}) {
    @card.Header() {
        @card.Title() { Dashboard Stats }
    }
    @card.Content() {
        <p class="text-dashboard-success">All systems operational</p>
    }
}
```
