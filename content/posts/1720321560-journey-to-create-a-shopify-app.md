# Journey to Create A Shopify App

When I started, I struggled with what I was going to use. I wanted to use Rails but I couldn't find much
information on how to use Rails with an embedded app in Shopify. That has since been debunked.

Before it was debunked, I settled on using Remix because it was in the Shopify developer documentation and made it easy 
to get started. Really, this turned out to be a huge pain in the ass both in development and with deployment.
<!--more-->
After a break from working on the shopify app, it dawned on me that I could use the javascript polaris libaries with a React
application within Rails.

I put all my shopify app routes under the `/shopify` URL so I can distingish between what routes are shopify app and what routes
are my own.

<br />

{{< gist haydenk d248642fda2c49b0ba0b830c9d63d7dd routes.rb >}}


The odd duck was the login url. It is not affected by the prefix for the `ShopifyApp::Engine`. It took quite a bit of digging but
I ended up figuring out you can set the login url in the shopify app initializer.

<br />
{{< gist haydenk d248642fda2c49b0ba0b830c9d63d7dd shopify_app.rb >}}


The javascript setup in rails has two entrypoints, one for `shopify.js` that the shopify app will load and `application.js` which is the 
standard Turbo JS setup for rails.

I had to modify the `package.json` and `bun.config.js` to "compile" both entrypoints.

<br />
{{< gist haydenk d248642fda2c49b0ba0b830c9d63d7dd bun.config.js >}}


<br />
{{< gist haydenk d248642fda2c49b0ba0b830c9d63d7dd package.json >}}


Bun will work with both application.js and shopify.js while sass and postcss will handle both application.bootstrap.scss and shopify.scss. So, now I can use both polaris styles and polaris javascript with my shopify app and use something else for the rest of the application.

Next steps: I am working on using Grape API to handle the JSON API requests, whether it be requests for the shopify app or for this application itself. The tricky part here is that some of the shopify endpoints will require shopify shop, and possibly user, authorization for it to work correctly.

---

Post Date: 2024-07-06T22:06:00-05:00

Tags: development, shopify, ruby, rails
