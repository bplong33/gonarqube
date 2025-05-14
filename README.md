# Gonarqube

A **Go** package for So**narqube**. This package was written out of personal
need for simplifying and automating administrative tasking in multiple
Sonarqube instances.

## Getting Started

...

## ToDo

- [ ] Add Validation for inputs
- [ ] Convert client params to accept all parameters rather than assume
      that the `url.Values` params passed by the user are correct
  > [!NOTE]
  > Current, the user is expected to build the `url.Values{}` and pass
  > them to the client functions. Ideally, the user should be able to look
  > at the function help and determine what options are accepted.
