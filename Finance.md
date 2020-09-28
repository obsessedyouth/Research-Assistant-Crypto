# Finance

1. Write a 500-word explanation of Bitcoin stock-to-flow model and make an argument for why it is a bad model?

## answer

I wrote a much longer essay that goes up to about
[1,500 plus words on Notion](https://www.notion.so/Unconventional-views-against-scarcity-73613a3ff10449e589a9d2b7ff595492)
you can check out the original form if interested.

## The argument against stock-to-flow

Value predates the existence of a market. Bitcoin is a currency, a different kind but still a currency that
by itself has no **intrinsic value.**

Much like gold and silver, or the Naira? Where it maps to reality is in the relationship to its **perceived value.**
Why does the price of bitcoin change? Micro economics tells you it's a function of a demand-supply curve and some
"utility function" but this is a fancy way of saying we all collectively came together to decide this and the
multivariate factors that affect our perception of how this thing maps to reality. The stock-to-flow model argues
that the price of bitcoin is driven by a constrained supply which reinforces this perception of value. I would like to
argue it's **how market participants use bitcoin that affects its price.**

"Use" here isn't just market speculation but the actual economic drive it propagates. Digital businesses, silk road 4.0?,
coin exchanges, cross-border transfers. It's how people use bitcoin, and there are all sorts of interesting reasons why.
Then the question becomes where is this information? and how can it tell us about the future?
It's hidden in the data and financial jargon. You have to go looking for how micro economic activity shows up in the
price model which the market naturally adjusts and responds to but not perfectly. This isn't necessarily scarcity,
as it is determined by the nature of bitcoin itself but **necessary scarcity as to how it affects an economic landscape.**
What creates such necessary scarcity? `Demand`. Demand and supply feed off of each other but demand must be created first
for the price of _anything_ to be percieved as valuable. The stock-flow model is useful for understanding the miner halving
effect on bitcoins price but it doesn't factor into it's variables if anyone would want those coins in the first place.
Models are not perfect, they expose some general truth provided certain assumptions hold true and in the real world the
stock-to-flow model is inadequate in factoring the demand side of the equation.

2. Yara Inc is listed on the NYSE with a stock price of $40 - the company is not known
 to pay dividends. We need to price a call option with a strike of $45 maturing in 4 months.
   The continuously-compounded risk-free rate is 3%/year, the mean return on the
   stock is 7%/year, and the standard deviation of the stock return is 40%/year. What is the Black-Scholes call price?

## answer

You can find details in `Finanace.pdf` which is separately rendered using `LaTex` or you can find the `tex` code
[here]().

### What is a call option

"The simplest kind of option is one that gives the right to buy a single share of common stock".
So what are we looking for here is simply the ability to buy/sell
\$40 dollars of Yara Inc. worth of stock within some condition. Given that this is listed on the `NYSE` it's safe to
infer this is an **American option** and this option can be exercised at any time until maturity. To earn this "ability"
you agree to pay some money with interest at some date if you don't exercise this right, this date being 4 months.

### What is the price of said option? What is a strike?

This is what you actually pay when said option right is exercised before the option matures. But why buy it in the first
place? Well stock price changes and sometimes? the value of the option is not equivalent to the stock price. Which is to
say that _sometimes_ the _ability_ to buy stock of Yara at $45 **may or may or not reflect the actual price of Yara stock at any
given moment in time.** This is how profit or loss is created. So depending on the date **the call option changes** far
options have cheaper costs and near options are expensive. If Yara rises to 46 $ yay! you made one dollar. If Yara never
goes beyond \$45 dollars? :( you lost whatever difference in price there.

### How can math help us make better decisions?

Well this can be represented by a differential equation. You have some quantity, in this case price, and you want to
understand how it behaves given some constraints within a specific period of time. The problem would be easier to solve
if we could reduce this phenomena to a _single dependent_ of the form _price = some constraint_ with respect
to time. The trouble arrives because we have **multiple dependent variables** of the form **price and volatility = some constraint**
wrt to time. The finance guys used to use this weird looking thing `C(S, T) = e^pT * SN(d_1)-(1-A)KN(d_2)` you can see the
[history of its development](http://www.ericbenhamou.net/documents/Encyclo/Pre%20Black-Scholes.pdf) the problem was they
had these magic constraint variables that were not accounted for, this was an algebraic equation that took natural
logarithmic. Math is hard. Especially when it's creative research math. You're trying to model some real world phenomena by
reducing it to first principles and creating an equation to describe said thing, it's an iterative kind of activity,
you're probably not gonna get it right the first few times you try. There's also other the ambiguity of "risk" to think
about.

### Introducing Black-Scholes Why a partial derivative?

Well this is actually based on a graph of real data, some other guys took some financial data, looked at the graph, drew
a big ol' curved line and made a formula out of it. The real insight? is this `equillibrium price = riskless asset`.
What does that mean exactly? The paper goes into how it's derived, I'm not brave enough quite yet lolol, but you can
check out the pdf for details.

### How to solve American style options?

This is black-scholes for finding option price on an European style call option:

![black-scholes european](https://wikimedia.org/api/rest_v1/media/math/render/svg/d85601f6192ee85748c2deef28240275510d634e)

The problem is the stock-exchange. Neither [NASDAQ](https://www.nasdaq.com/articles/option-styles-2019-06-10) nor
[NYSE](https://www.nyse.com/products/options-equity) offer European style call options. So how do we apply the
Black-Scholes equation, to American style options?
[apparently we have to use an Extended version by someone called Merton](https://money.stackexchange.com/questions/8898/does-the-black-scholes-model-apply-to-american-style-options)

[Well, there's cutting edge research but sadly it's behind a pay-wall](https://papers.ssrn.com/sol3/papers.cfm?abstract_id=1516172)

Luckily, this is a theoretical question. For an ideal scenario of no dividend `American call option == European call option`.

To solve the linear call price it is given by `C(S, T) = SN(x_{1}) - BN(x_{2})` which is equal to `$2.02361738929`
