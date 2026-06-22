#!/usr/bin/env python3

import aws_cdk as cdk

from data_enrichment_platform.data_enrichment_platform_stack import DataEnrichmentPlatformStack


app = cdk.App()
DataEnrichmentPlatformStack(app, "DataEnrichmentPlatformStack")

app.synth()
