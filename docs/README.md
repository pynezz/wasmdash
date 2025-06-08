# Documentation

## Table of contents

- [Theme middleware](#theme-middleware)
- [Widgets](#widgets)

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

## Widgets

Widgets are reusable components that can be used to build complex user interfaces. They are defined in the `pkg/components/widgets` package.

### Adding a widget

> [!TIP]
> Planned feature

Adding a widget is a planned feature, outlined as follows:

1. Define the widget in a "widgets/WIDGET_NAME" directory.
2. widget.toml
 - the widget is parsed by the widget parser
 - the widget is built and rendered by templ

## Dynamic layout

> [!TIP]
> Planned feature

Dynamic layout is not about responsive design, but about the ability to dynamically change the layout of a page based on the data it contains, or events it receives. One example is that if you've defined a widget with a showOnEvent attribute,
the widget will be shown when the event is triggered.

First planned natively supported dynamic widget is a lightning radar widget.

It'll not be visible until a lightning strike is detected within a certain radius.
