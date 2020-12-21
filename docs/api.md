# API

## Batch Create Occurrences

Add multiple occurrences

### gRPC

BatchCreateOccurrences([BatchCreateOccurrencesRequest](grpc.md#rode.v1alpha1.BatchCreateOccurrencesRequest)) [BatchCreateOccurrencesResponse](grpc.md#rode.v1alpha1.BatchCreateOccurrencesResponse)

### REST

TODO

---

## Attest Policy

**Work in progress**

Verify that an artifact satisfies a policy.

### gRPC

**AttestPolicy**([AttestPolicyRequest](grpc.md#rode.v1alpha1.AttestPolicyRequest)) [AttestPolicyResponse](grpc.md#rode.v1alpha1.AttestPolicyResponse)

### REST


`GET /api/attest-policy`


**Request**
```json
{
  "policy": "my-policy",
  "resourceURI": "https://my-repository.org/project/artifact
}
```

**Response**

The policy attestation response indicates the current state of the policy for the artifact (pass, fail), wether the policy state has changed since it was last evaluated, and the history of attestations and violations.

```json
{
  "allow": false,
  "changed": true,
  "attestations" : [
    {
      "allow": false,
      "datetime": "Tue Nov 24 07:50:16 PST 2020",
      "violations": [
        {
          "id": "sonarqube_qualitygate_fail",
          "name": "SonarQube Quality Gate Failed",
          "description": "A SonarQube quality gate failed to meet one of its conditions. Please see addition violations for more information.",
          "link": "https://sonarqube.my.org/scandetails",
        },
        {
          "id": "sonarqube_qualitygate_condition_coverage",
          "name": "SonarQube Code Coverage Condition Failed",
          "description": "A SonarQube quality gate condition failed because it did not meet Code Coverage minimum",
          "link": "https://sonarqube.my.org/scandetails",
        }
        ...
      ]
    }
  ],
  ...
}
