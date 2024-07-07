---
layout: post
title: Journey to Create A Shopify App
date: 2024-07-06 22:06 -0500
---

When I started, I struggled with what I was going to use. I wanted to use Rails but I couldn't find much
information on how to use Rails with an embedded app in Shopify. That has since been debunked.

Before it was debunked, I settled on using Remix because it was in the Shopify developer documentation and made it easy 
to get started. Really, this turned out to be a huge pain in the ass both in development and with deployment.
<!--more-->
After a break from working on the shopify app, it dawned on me that I could use the javascript polaris libaries with a React
application within Rails.

I put all my shopify app routes under the `/shopify` URL so I can distingish between what routes are shopify app and what routes
are my own.

`config/routes.rb`
```ruby 
  get '/shopify/', :to => 'shopify#index'
  mount ShopifyApp::Engine, at: '/shopify'
```

The odd duck was the login url. It is not affected by the prefix for the `ShopifyApp::Engine`. It took quite a bit of digging but
I ended up figuring out you can set the login url in the shopify app initializer.

`config/initializers/shopify_app.rb`
```ruby
ShopifyApp.configure do |config|
  ...
  config.login_url = '/shopify/login'
  ...
end
```

The javascript setup in rails has two entrypoints, one for `shopify.js` that the shopify app will load and `application.js` which is the 
standard Turbo JS setup for rails.

I had to modify the `package.json` and `bun.config.js` to "compile" both entrypoints.

`bun.config.js`
```js
const config = {
  sourcemap: "external",
  entrypoints: ["app/javascript/application.js", "app/javascript/shopify.js"],
  outdir: path.join(process.cwd(), "app/assets/builds"),
};
```

`package.json`
```json
{
  "scripts": {
    "build:css:compile": "sass ./app/assets/stylesheets/application.bootstrap.scss:./app/assets/builds/application.css ./app/assets/stylesheets/shopify.scss:./app/assets/builds/shopify.css --no-source-map --load-path=node_modules",
    "build:css:prefix": "postcss ./app/assets/builds/*.css --use=autoprefixer --dir=./app/assets/builds",
    "build:css": "bun run build:css:compile && bun run build:css:prefix",
    "watch:css": "nodemon --watch ./app/assets/stylesheets/ --ext scss --exec \"bun run build:css\"",
    "build": "bun bun.config.js"
  },
}
```

Bun will work with both application.js and shopify.js while sass and postcss will handle both application.bootstrap.scss and shopify.scss. So, now I can use both polaris styles and polaris javascript with my shopify app and use something else for the rest of the application.

Next steps: I am working on using Grape API to handle the JSON API requests, whether it be requests for the shopify app or for this application itself. The tricky part here is that some of the shopify endpoints will require shopify shop, and possibly user, authorization for it to work correctly.
