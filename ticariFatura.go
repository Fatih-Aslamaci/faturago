package main

var (
	xmlTest2 = `
<person xmlns="ns1" xmlns:ns2="ns2">
  <name>Oliver</name>
  <ns2:phone>110</ns2:phone>
</person>
	`

	xmlTest = `
<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet type="text/xsl" href="general.xslt"?>
<Invoice 
xmlns="urn:oasis:names:specification:ubl:schema:xsd:Invoice-2" 
xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
xmlns:cac="urn:oasis:names:specification:ubl:schema:xsd:CommonAggregateComponents-2" 
xmlns:ext="urn:oasis:names:specification:ubl:schema:xsd:CommonExtensionComponents-2" 
xmlns:ds="http://www.w3.org/2000/09/xmldsig#" 
xmlns:xades="http://uri.etsi.org/01903/v1.3.2#" 
xmlns:cbc="urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2" 
xsi:schemaLocation="urn:oasis:names:specification:ubl:schema:xsd:Invoice-2 ..\xsdrt\maindoc\UBL-Invoice-2.1.xsd" 
xmlns:n4="http://www.altova.com/samplexml/other-namespace" >
	<ext:UBLExtensions>
		<ext:UBLExtension>
			<ext:ExtensionContent>
				<n4:auto-generated_for_wildcard/>
			</ext:ExtensionContent>
		</ext:UBLExtension>
	</ext:UBLExtensions>
	<cbc:UBLVersionID>2.1</cbc:UBLVersionID>
	<cbc:CustomizationID>TR1.2</cbc:CustomizationID>
	<cbc:ProfileID>TICARIFATURA</cbc:ProfileID>
	<cbc:ID>GIB2009000000011</cbc:ID>
	<cbc:CopyIndicator>false</cbc:CopyIndicator>
	<cbc:UUID>F47AC10B-58CC-4372-A567-0E02B2C3D479</cbc:UUID>
	<cbc:IssueDate>2009-01-05</cbc:IssueDate>
	<cbc:IssueTime>14:42:00</cbc:IssueTime>
	<cbc:InvoiceTypeCode>SATIS</cbc:InvoiceTypeCode>
	<cbc:DocumentCurrencyCode>TRY</cbc:DocumentCurrencyCode>
	<cbc:LineCountNumeric>8</cbc:LineCountNumeric>
	<cac:DespatchDocumentReference>
		<cbc:ID>180921</cbc:ID>
		<cbc:IssueDate>2009-01-02</cbc:IssueDate>
	</cac:DespatchDocumentReference>
	<cac:Signature>
		<cbc:ID schemeID="VKN_TCKN">1288331521</cbc:ID>
		<cac:SignatoryParty>
			<cac:PartyIdentification>
				<cbc:ID schemeID="VKN">1288331521</cbc:ID>
			</cac:PartyIdentification>
			<cac:PostalAddress>
				<cbc:StreetName>Papatya Caddesi Yasemin Sokak</cbc:StreetName>
				<cbc:BuildingNumber>21</cbc:BuildingNumber>
				<cbc:CitySubdivisionName>Beşiktaş</cbc:CitySubdivisionName>
				<cbc:CityName>İstanbul</cbc:CityName>
				<cbc:PostalZone>34100</cbc:PostalZone>
				<cac:Country>
					<cbc:Name>Türkiye</cbc:Name>
				</cac:Country>
			</cac:PostalAddress>
		</cac:SignatoryParty>
		<cac:DigitalSignatureAttachment>
			<cac:ExternalReference>
				<cbc:URI>#Signature</cbc:URI>
			</cac:ExternalReference>
		</cac:DigitalSignatureAttachment>
	</cac:Signature>
 	<cac:AccountingSupplierParty>
		<cac:Party>
			<cbc:WebsiteURI>http://www.aaa.com.tr/</cbc:WebsiteURI>
			<cac:PartyIdentification>
				<cbc:ID schemeID="VKN">1288331521</cbc:ID>
			</cac:PartyIdentification>
			<cac:PartyName>
				<cbc:Name>AAA Anonim Şirketi</cbc:Name>
			</cac:PartyName>
			<cac:PostalAddress>
				<cbc:ID>1234567890</cbc:ID>
				<cbc:StreetName>Papatya Caddesi Yasemin Sokak</cbc:StreetName>
				<cbc:BuildingNumber>21</cbc:BuildingNumber>
				<cbc:CitySubdivisionName>Beşiktaş</cbc:CitySubdivisionName>
				<cbc:CityName>İstanbul</cbc:CityName>
				<cbc:PostalZone>34100</cbc:PostalZone>
				<cac:Country>
					<cbc:Name>Türkiye</cbc:Name>
				</cac:Country>
			</cac:PostalAddress>
			<cac:PartyTaxScheme>
				<cac:TaxScheme>
					<cbc:Name>Büyük Mükellefler</cbc:Name>
				</cac:TaxScheme>
			</cac:PartyTaxScheme>
			<cac:Contact>
				<cbc:Telephone>(212) 925 51515</cbc:Telephone>
				<cbc:Telefax>(212) 925505015</cbc:Telefax>
				<cbc:ElectronicMail>aa@aaa.com.tr</cbc:ElectronicMail>
			</cac:Contact>
		</cac:Party>
	</cac:AccountingSupplierParty>
	<cac:AccountingCustomerParty>
		<cac:Party>
			<cbc:WebsiteURI>http://www.bbb.com.tr/</cbc:WebsiteURI>
			<cac:PartyIdentification>
				<cbc:ID schemeID="VKN">9205121120</cbc:ID>
			</cac:PartyIdentification>
			<cac:PartyName>
				<cbc:Name>BBB Limited  Şirketi</cbc:Name>
			</cac:PartyName>
			<cac:PostalAddress>
				<cbc:ID>1234567890</cbc:ID>
				<cbc:StreetName>Ihlamur Mahallesi Selvi Caddesi Sedir Sokak</cbc:StreetName>
				<cbc:BuildingNumber>75/A</cbc:BuildingNumber>
				<cbc:CitySubdivisionName>Kızılay</cbc:CitySubdivisionName>
				<cbc:CityName>Ankara</cbc:CityName>
				<cbc:PostalZone>06100</cbc:PostalZone>
				<cac:Country>
					<cbc:Name>Türkiye</cbc:Name>
				</cac:Country>
			</cac:PostalAddress>
			<cac:PartyTaxScheme>
				<cac:TaxScheme>
					<cbc:Name>Çankaya</cbc:Name>
				</cac:TaxScheme>
			</cac:PartyTaxScheme>
			<cac:Contact>
				<cbc:Telephone>(312) 621 1111</cbc:Telephone>
				<cbc:Telefax>(312) 621 1010</cbc:Telefax>
				<cbc:ElectronicMail>bb@bbb.com.tr</cbc:ElectronicMail>
			</cac:Contact>
		</cac:Party>
	</cac:AccountingCustomerParty>
	<cac:PaymentMeans>
		<cbc:PaymentMeansCode>1</cbc:PaymentMeansCode>
		<cac:PayeeFinancialAccount>
			<cbc:ID>5652214414</cbc:ID>
			<cbc:CurrencyCode>TRY</cbc:CurrencyCode>
			<cbc:PaymentNote>RRR Bankası Beşiktaş Şubesi TL Hesabı</cbc:PaymentNote>
		</cac:PayeeFinancialAccount>
	</cac:PaymentMeans>
	<cac:PaymentTerms>
		<cbc:Note>Fatura düzenlenme tarihinden itibaren 20 gün içerisinde ödenecektir.</cbc:Note>
		<cbc:PaymentDueDate>2008-11-25</cbc:PaymentDueDate>
	</cac:PaymentTerms>
	<cac:TaxTotal>
		<cbc:TaxAmount currencyID="TRY">4538.97</cbc:TaxAmount>
		<cac:TaxSubtotal>
			<cbc:TaxableAmount currencyID="TRY">25216.50</cbc:TaxableAmount>
			<cbc:TaxAmount currencyID="TRY">4538.97</cbc:TaxAmount>
			<cbc:CalculationSequenceNumeric>1.0</cbc:CalculationSequenceNumeric>
			<cbc:Percent>18</cbc:Percent>
			<cac:TaxCategory>
				<cac:TaxScheme>
					<cbc:Name>Katma Değer Vergisi</cbc:Name>
					<cbc:TaxTypeCode>0015</cbc:TaxTypeCode>
				</cac:TaxScheme>
			</cac:TaxCategory>
		</cac:TaxSubtotal>
	</cac:TaxTotal>
	<cac:LegalMonetaryTotal>
		<cbc:LineExtensionAmount currencyID="TRY">26003.40</cbc:LineExtensionAmount>
		<cbc:TaxExclusiveAmount currencyID="TRY">25216.50</cbc:TaxExclusiveAmount>
		<cbc:TaxInclusiveAmount currencyID="TRY">29755.47</cbc:TaxInclusiveAmount>
		<cbc:AllowanceTotalAmount currencyID="TRY">786.90 </cbc:AllowanceTotalAmount>
		<cbc:PayableAmount currencyID="TRY">29755.47</cbc:PayableAmount>
	</cac:LegalMonetaryTotal>
    <cac:InvoiceLine>
        <cbc:ID>1</cbc:ID>
        <cbc:InvoicedQuantity unitCode="NIU">12</cbc:InvoicedQuantity>
        <cbc:LineExtensionAmount currencyID="TRY">9906</cbc:LineExtensionAmount>
        <cac:AllowanceCharge>
            <cbc:ChargeIndicator>false</cbc:ChargeIndicator>
            <cbc:MultiplierFactorNumeric>0.05</cbc:MultiplierFactorNumeric>
            <cbc:Amount currencyID="TRY">495.3</cbc:Amount>
            <cbc:BaseAmount currencyID="TRY">9906</cbc:BaseAmount>
        </cac:AllowanceCharge>
        <cac:TaxTotal>
            <cbc:TaxAmount currencyID="TRY">1693.93</cbc:TaxAmount>
            <cac:TaxSubtotal>
                <cbc:TaxableAmount currencyID="TRY">9410.7</cbc:TaxableAmount>
                <cbc:TaxAmount currencyID="TRY">1693.93</cbc:TaxAmount>
                <cbc:Percent>18</cbc:Percent>
                <cac:TaxCategory>
                    <cac:TaxScheme>
                        <cbc:Name>KDV</cbc:Name>
                        <cbc:TaxTypeCode>0015</cbc:TaxTypeCode>
                    </cac:TaxScheme>
                </cac:TaxCategory>
            </cac:TaxSubtotal>
        </cac:TaxTotal>
        <cac:Item>
            <cbc:Name>Masa Üstü Bilgisayar</cbc:Name>
        </cac:Item>
        <cac:Price>
            <cbc:PriceAmount currencyID="TRY">825.5</cbc:PriceAmount>
        </cac:Price>
    </cac:InvoiceLine>
    <cac:InvoiceLine>
        <cbc:ID>2</cbc:ID>
        <cbc:InvoicedQuantity unitCode="NIU">8</cbc:InvoicedQuantity>
        <cbc:LineExtensionAmount currencyID="TRY">9720</cbc:LineExtensionAmount>
        <cac:AllowanceCharge>
            <cbc:ChargeIndicator>false</cbc:ChargeIndicator>
            <cbc:MultiplierFactorNumeric>0.03</cbc:MultiplierFactorNumeric>
            <cbc:Amount currencyID="TRY">291.6</cbc:Amount>
            <cbc:BaseAmount currencyID="TRY">9720</cbc:BaseAmount>
        </cac:AllowanceCharge>
        <cac:TaxTotal>
            <cbc:TaxAmount currencyID="TRY">1697.11</cbc:TaxAmount>
            <cac:TaxSubtotal>
                <cbc:TaxableAmount currencyID="TRY">9428.4</cbc:TaxableAmount>
                <cbc:TaxAmount currencyID="TRY">1697.11</cbc:TaxAmount>
                <cbc:Percent>18</cbc:Percent>
                <cac:TaxCategory>
                    <cac:TaxScheme>
                        <cbc:Name>KDV</cbc:Name>
                        <cbc:TaxTypeCode>0015</cbc:TaxTypeCode>
                    </cac:TaxScheme>
                </cac:TaxCategory>
            </cac:TaxSubtotal>
        </cac:TaxTotal>
        <cac:Item>
            <cbc:Name>Notebook Bilgisayar</cbc:Name>
        </cac:Item>
        <cac:Price>
            <cbc:PriceAmount currencyID="TRY">1215</cbc:PriceAmount>
        </cac:Price>
    </cac:InvoiceLine>
    <cac:InvoiceLine>
        <cbc:ID>3</cbc:ID>
        <cbc:InvoicedQuantity unitCode="NIU">18</cbc:InvoicedQuantity>
        <cbc:LineExtensionAmount currencyID="TRY">688.5</cbc:LineExtensionAmount>
        <cac:AllowanceCharge>
            <cbc:ChargeIndicator>false</cbc:ChargeIndicator>
            <cbc:MultiplierFactorNumeric>0.0</cbc:MultiplierFactorNumeric>
            <cbc:Amount currencyID="TRY">0</cbc:Amount>
            <cbc:BaseAmount currencyID="TRY">688.5</cbc:BaseAmount>
        </cac:AllowanceCharge>
        <cac:TaxTotal>
            <cbc:TaxAmount currencyID="TRY">123.93</cbc:TaxAmount>
            <cac:TaxSubtotal>
                <cbc:TaxableAmount currencyID="TRY">688.5</cbc:TaxableAmount>
                <cbc:TaxAmount currencyID="TRY">123.93</cbc:TaxAmount>
                <cbc:Percent>18</cbc:Percent>
                <cac:TaxCategory>
                    <cac:TaxScheme>
                        <cbc:Name>KDV</cbc:Name>
                        <cbc:TaxTypeCode>0015</cbc:TaxTypeCode>
                    </cac:TaxScheme>
                </cac:TaxCategory>
            </cac:TaxSubtotal>
        </cac:TaxTotal>
        <cac:Item>
            <cbc:Name>Notebook Çantası</cbc:Name>
        </cac:Item>
        <cac:Price>
            <cbc:PriceAmount currencyID="TRY">38.25</cbc:PriceAmount>
        </cac:Price>
    </cac:InvoiceLine>
    <cac:InvoiceLine>
        <cbc:ID>4</cbc:ID>
        <cbc:InvoicedQuantity unitCode="NIU">13</cbc:InvoicedQuantity>
        <cbc:LineExtensionAmount currencyID="TRY">1576.9</cbc:LineExtensionAmount>
        <cac:AllowanceCharge>
            <cbc:ChargeIndicator>false</cbc:ChargeIndicator>
            <cbc:MultiplierFactorNumeric>0.0</cbc:MultiplierFactorNumeric>
            <cbc:Amount currencyID="TRY">0</cbc:Amount>
            <cbc:BaseAmount currencyID="TRY">1576.9</cbc:BaseAmount>
        </cac:AllowanceCharge>
        <cac:TaxTotal>
            <cbc:TaxAmount currencyID="TRY">283.84</cbc:TaxAmount>
            <cac:TaxSubtotal>
                <cbc:TaxableAmount currencyID="TRY">1576.9</cbc:TaxableAmount>
                <cbc:TaxAmount currencyID="TRY">283.84</cbc:TaxAmount>
                <cbc:Percent>18</cbc:Percent>
                <cac:TaxCategory>
                    <cac:TaxScheme>
                        <cbc:Name>KDV</cbc:Name>
                        <cbc:TaxTypeCode>0015</cbc:TaxTypeCode>
                    </cac:TaxScheme>
                </cac:TaxCategory>
            </cac:TaxSubtotal>
        </cac:TaxTotal>
        <cac:Item>
            <cbc:Name>Yazıcı </cbc:Name>
        </cac:Item>
        <cac:Price>
            <cbc:PriceAmount currencyID="TRY">121.3</cbc:PriceAmount>
        </cac:Price>
    </cac:InvoiceLine>
    <cac:InvoiceLine>
        <cbc:ID>5</cbc:ID>
        <cbc:InvoicedQuantity unitCode="BX">200</cbc:InvoicedQuantity>
        <cbc:LineExtensionAmount currencyID="TRY">1400</cbc:LineExtensionAmount>
        <cac:AllowanceCharge>
            <cbc:ChargeIndicator>false</cbc:ChargeIndicator>
            <cbc:MultiplierFactorNumeric>0.0</cbc:MultiplierFactorNumeric>
            <cbc:Amount currencyID="TRY">0</cbc:Amount>
            <cbc:BaseAmount currencyID="TRY">1400</cbc:BaseAmount>
        </cac:AllowanceCharge>
        <cac:TaxTotal>
            <cbc:TaxAmount currencyID="TRY">252</cbc:TaxAmount>
            <cac:TaxSubtotal>
                <cbc:TaxableAmount currencyID="TRY">1400</cbc:TaxableAmount>
                <cbc:TaxAmount currencyID="TRY">252</cbc:TaxAmount>
                <cbc:Percent>18</cbc:Percent>
                <cac:TaxCategory>
                    <cac:TaxScheme>
                        <cbc:Name>KDV</cbc:Name>
                        <cbc:TaxTypeCode>0015</cbc:TaxTypeCode>
                    </cac:TaxScheme>
                </cac:TaxCategory>
            </cac:TaxSubtotal>
        </cac:TaxTotal>
        <cac:Item>
            <cbc:Name>CD-R 2*56</cbc:Name>
        </cac:Item>
        <cac:Price>
            <cbc:PriceAmount currencyID="TRY">7</cbc:PriceAmount>
        </cac:Price>
    </cac:InvoiceLine>
    <cac:InvoiceLine>
        <cbc:ID>6</cbc:ID>
        <cbc:InvoicedQuantity unitCode="BX">200</cbc:InvoicedQuantity>
        <cbc:LineExtensionAmount currencyID="TRY">2200</cbc:LineExtensionAmount>
        <cac:AllowanceCharge>
            <cbc:ChargeIndicator>false</cbc:ChargeIndicator>
            <cbc:MultiplierFactorNumeric>0.0</cbc:MultiplierFactorNumeric>
            <cbc:Amount currencyID="TRY">0</cbc:Amount>
            <cbc:BaseAmount currencyID="TRY">2200</cbc:BaseAmount>
        </cac:AllowanceCharge>
        <cac:TaxTotal>
            <cbc:TaxAmount currencyID="TRY">396</cbc:TaxAmount>
            <cac:TaxSubtotal>
                <cbc:TaxableAmount currencyID="TRY">2200</cbc:TaxableAmount>
                <cbc:TaxAmount currencyID="TRY">396</cbc:TaxAmount>
                <cbc:Percent>18</cbc:Percent>
                <cac:TaxCategory>
                    <cac:TaxScheme>
                        <cbc:Name>KDV</cbc:Name>
                        <cbc:TaxTypeCode>0015</cbc:TaxTypeCode>
                    </cac:TaxScheme>
                </cac:TaxCategory>
            </cac:TaxSubtotal>
        </cac:TaxTotal>
        <cac:Item>
            <cbc:Name>DVD-R</cbc:Name>
        </cac:Item>
        <cac:Price>
            <cbc:PriceAmount currencyID="TRY">11</cbc:PriceAmount>
        </cac:Price>
    </cac:InvoiceLine>
    <cac:InvoiceLine>
        <cbc:ID>7</cbc:ID>
        <cbc:InvoicedQuantity unitCode="NIU">80</cbc:InvoicedQuantity>
        <cbc:LineExtensionAmount currencyID="TRY">340</cbc:LineExtensionAmount>
        <cac:AllowanceCharge>
            <cbc:ChargeIndicator>false</cbc:ChargeIndicator>
            <cbc:MultiplierFactorNumeric>0.0</cbc:MultiplierFactorNumeric>
            <cbc:Amount currencyID="TRY">0</cbc:Amount>
            <cbc:BaseAmount currencyID="TRY">340</cbc:BaseAmount>
        </cac:AllowanceCharge>
        <cac:TaxTotal>
            <cbc:TaxAmount currencyID="TRY">61.2</cbc:TaxAmount>
            <cac:TaxSubtotal>
                <cbc:TaxableAmount currencyID="TRY">340</cbc:TaxableAmount>
                <cbc:TaxAmount currencyID="TRY">61.2</cbc:TaxAmount>
                <cbc:Percent>18</cbc:Percent>
                <cac:TaxCategory>
                    <cac:TaxScheme>
                        <cbc:Name>KDV</cbc:Name>
                        <cbc:TaxTypeCode>0015</cbc:TaxTypeCode>
                    </cac:TaxScheme>
                </cac:TaxCategory>
            </cac:TaxSubtotal>
        </cac:TaxTotal>
        <cac:Item>
            <cbc:Name>Dolma Kalem</cbc:Name>
        </cac:Item>
        <cac:Price>
            <cbc:PriceAmount currencyID="TRY">4.25</cbc:PriceAmount>
        </cac:Price>
    </cac:InvoiceLine>
    <cac:InvoiceLine>
        <cbc:ID>8</cbc:ID>
        <cbc:InvoicedQuantity unitCode="NIU">80</cbc:InvoicedQuantity>
        <cbc:LineExtensionAmount currencyID="TRY">172</cbc:LineExtensionAmount>
        <cac:AllowanceCharge>
            <cbc:ChargeIndicator>false</cbc:ChargeIndicator>
            <cbc:MultiplierFactorNumeric>0.0</cbc:MultiplierFactorNumeric>
            <cbc:Amount currencyID="TRY">0</cbc:Amount>
            <cbc:BaseAmount currencyID="TRY">172</cbc:BaseAmount>
        </cac:AllowanceCharge>
        <cac:TaxTotal>
            <cbc:TaxAmount currencyID="TRY">30.96</cbc:TaxAmount>
            <cac:TaxSubtotal>
                <cbc:TaxableAmount currencyID="TRY">172</cbc:TaxableAmount>
                <cbc:TaxAmount currencyID="TRY">30.96</cbc:TaxAmount>
                <cbc:Percent>18</cbc:Percent>
                <cac:TaxCategory>
                    <cac:TaxScheme>
                        <cbc:Name>KDV</cbc:Name>
                        <cbc:TaxTypeCode>0015</cbc:TaxTypeCode>
                    </cac:TaxScheme>
                </cac:TaxCategory>
            </cac:TaxSubtotal>
        </cac:TaxTotal>
        <cac:Item>
            <cbc:Name>Tükenmez Kalem</cbc:Name>
        </cac:Item>
        <cac:Price>
            <cbc:PriceAmount currencyID="TRY">2.15</cbc:PriceAmount>
        </cac:Price>
    </cac:InvoiceLine>
</Invoice>
`
)
