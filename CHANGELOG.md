# develop

- Add `client.AddWiki`.
- Add `client.UpdateWiki`.
- Add `client.DeleteWiki`.
- Add `client.GetWikiCount`.
- Add `client.GetWikiTags`.
- Fix test.
- Add `client.GetWikiAttachments`.
- Add `client.CreateWikiAttachment`.
- Add `client.DeleteWikiAttachment`.
- Add `client.DownloadWikiAttachment`.

## v0.4.1

- Fix error message.

## v0.4.0

- Add error variables.
- `types.Errors` implements error interface.

## v0.3.0

- Change method arguments of `client.Client.GetAllIssues`.
- Remove unnecessary parameter of `client.Client.GetPullRequestsCount`.
- Fix method arguments of `client.Client.GetPullRequestComments`.
- Fix some types in `types` package.

## v0.2.0

- While debugging, fix panic when request body is nil.
- Add `EncodeQuery` and `Uniq` methods to `types.Issue`.
- Change method signature of `client.Client.AddIssue`.
- Change method signature of `client.Client.UpdateIssue`.
- Fix types.
- Remove unnecessary query value.
- Add `client.Client.GetAllIssues` method.
- Add `client.Client.GetAllPullRequests` method.

## v0.1.0

- Rename function `client.NewClient` to `client.New`.
- Fix some types
- `Date` supports nil value.
- Add `types.Hours`.
- Change parameter of `Client.AddIssue`.
- Add log messages about payload.

## v0.0.0

Initial version
