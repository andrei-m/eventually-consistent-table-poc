# eventually-consistent-table-poc

**WORK IN PROGRESS**

Good UX for presenting a table sourced from multiple eventually-consistent sources

# Development

```
go install
npm install

# Start the backend first
$GOPATH/bin/eventually-consistent-table-poc

# Start the frontend second
npm run dev
```

## Problem

Complex applications often need to present data stitched together from multiple sources. For example, an application primarily responsible for CRUD operations may evolve to present metrics associated with created entities. Those metrics may be sourced from another system, such as an analytics database. A normal technical design constraint within this architecture is that the two systems in play are not strongly consistent: the analytics database is often 'behind' the entity (aka transactional) database, which puts burden on the downstream application to handle the potential inconsistency gracefully.

  * The system should be generally usable even if the eventually-consistent system is behind. Replication latency to the eventually-consistent system should not prevent functionality that can be enabled solely by referencing a separate system that is not latent.
  * The stitching-together should be abstracted from user perception as much as possible. Users should be able to operate on the stitched together data as if it was one consolidated record, even if the implementation is more complex. For example, stitched-together data presented in a table should allow for sorting and filtering controls on columns sourced from either the transactional or eventually-consistent data source in a seamless way.

For this proof of concept, the problem is limited to presenting stitched-together data in a table. Each row can contain data from a transactional system, and an eventually-consistent follower system.

## Inspirations

### Data Integration into the transactional database
This approach brings eventually-consistent data back into the data source that serves CRUD operations, which is is more often strongly-consistent. This makes the application's job easy at the expense of significant additional responsibility for the CRUD database. For example, it may have to take on analytics-like queries if analytics workloads are added to its overall responsibility.

### Shasta
Google's [Shasta paper](https://static.googleusercontent.com/media/research.google.com/en//pubs/archive/45394.pdf), originally published in 2016 describes a best-of-all-worlds architecture describing a myriad of complex data sources with few guarantees, low response time, low data replication latency, and complete abstraction of the stitching-together through a middleware layer. This middleware layer is implies that user-perceptible front-ends do not present data piecemeal. Instead, the middleware presents stitched-together records following a server-side data integration. This approach is impressive, but requires a significant investment in engineering and systems to achieve at scale.

### Front-End Data Integration
The user-perceptible behavior of the [AWS IAM user list page](https://us-east-1.console.aws.amazon.com/iamv2/home?region=us-east-1#/users) (ostensibly "aimv2" based on the URL) demonstrates a front-end join approach to this problem. On an initial page load or refresh, it is clear that various components of the list, such as each user's group memberships, last activity time, password and access key ages, etc. load piecemeal following an initial page render. If a user list contains more users than the current page size, sorting by any of these fields causes a cascade of partial reloads, as the users table is reconstructed to match the new requested sorting criteria. It is clear that this is achieved through multiple requests to services, with data stitched together on the front-end. Despite some user-perceptible impact of this solution to the eventual consistency problem (certain columns remain in 'loading' state while the stitching happens on the front-end), the page remains largely pleasant to use:

  * The primary sorting column and page number (or [pagination cursor](https://ignaciochiazzo.medium.com/paginating-requests-in-apis-d4883d4c1c4c)) determine the entities presented on the page. This serves as the 'backbone' from which all other rendered data are requested. Rendering this first unblocks a lot of usability while the rest of the page renders, particularly when the sort-by column is a reference to an entity field from a source-of-truth data source (i.e. there is no need to wait for eventually consistent data to unblock functionality available solely from the source-of-truth).
  * A large part of the UX pleasantness comes from low response time from each of the subsequent service calls. These are individually easier to optimize than a system that strives to optimize for low response time for a complete stitched-together result (bound by the slowest eventually consistent service).

## Solution Approach

This proof of concept applies learnings the Front-End Data Integration approach.

  1. A sort-by column, pagination cursor, and page size are used to establish the set of stitched-together records presented in a tabular view. This is referred to as the 'backbone' for the rest of page rendering.
  2. The identifiers from the backbone are used to query for column values from other data sources individually (called 'appendages'). A front-end join performs the data stitching.

Graceful degradation techniques are applied to handle latent data sources, including the backbone:

  * When sufficient data is present in the backbone, but an appendage is missing data, sensible defaults are rendered in place of the latent data
  * When the backbone is based on an eventually consistent data set and pagination parameters leave open the potential for records not yet integrated into the backbone, the backbone is augmented with data gleaned from appendages (including the source of truth)

