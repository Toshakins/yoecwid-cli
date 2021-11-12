package main

const storeBaseTemplate = `<?xml version="1.0" ?>
<template>
	<owner>
		<companyEmail>%s</companyEmail>
		<fromEmail>%s</fromEmail>
		<fromAdminEmail>%s</fromAdminEmail>
		<storeConfiguration>
			<source>YOLA</source>
			<currency>USD</currency>
			<currencyPrefix>$</currencyPrefix>
			<paymentmethods>
				<paymentmethod>
					<enabled>false</enabled>
					<instructions/>
					<name>Credit card</name>
				</paymentmethod>
				<paymentmethod>
					<enabled>false</enabled>
					<instructions/>
					<name>Echeck</name>
				</paymentmethod>
				<paymentmethod>
					<enabled>false</enabled>
					<instructions/>
					<name>PayPal</name>
					<processor>paypalStandard</processor>
				</paymentmethod>
				<paymentmethod>
					<enabled>false</enabled>
					<instructions/>
					<name>Check</name>
					<processor>offlineCheck</processor>
				</paymentmethod>
				<paymentmethod>
					<enabled>false</enabled>
					<instructions/>
					<name>Business check</name>
					<processor>offline</processor>
				</paymentmethod>
				<paymentmethod>
					<enabled>false</enabled>
					<instructions/>
					<name>Fax order</name>
					<processor>offline</processor>
				</paymentmethod>
				<paymentmethod>
					<enabled>false</enabled>
					<instructions/>
					<name>Money order</name>
					<processor>offline</processor>
				</paymentmethod>
				<paymentmethod>
					<enabled>false</enabled>
					<instructions/>
					<name>Personal check</name>
					<processor>offline</processor>
				</paymentmethod>
				<paymentmethod>
					<enabled>false</enabled>
					<instructions/>
					<name>Phone order</name>
					<processor>offline</processor>
				</paymentmethod>
				<paymentmethod>
					<enabled>false</enabled>
					<instructions/>
					<name>Purchase order</name>
					<processor>offlinePurchaseOrder</processor>
				</paymentmethod>
				<paymentmethod>
					<enabled>false</enabled>
					<instructions/>
					<name>Wire transfer</name>
					<paymentInstructionsTitle>Our bank account details</paymentInstructionsTitle>
					<processor>offline</processor>
				</paymentmethod>
			</paymentmethods>
			<shippingAndTaxSettings/>
		</storeConfiguration>
	</owner>
	<categories/>
	<products>
		<product id="0000">
			<description>This is a product description. Add more detail about this product, such as benefits, appearance, components and value</description>
			<name>Product 1</name>
			<price>199.00</price>
			<sku>0000</sku>
			<image>
				<link>
					<imageLink>
						<url>https://assets.yolacdn.net/store/generic/image-0.jpg</url>
					</imageLink>
				</link>
			</image>
		</product>
		<product id="0001">
			<description>This is a product description. Add more detail about this product, such as benefits, appearance, components and value</description>
			<name>Product 2</name>
			<price>29.00</price>
			<sku>0001</sku>
			<image>
				<link>
					<imageLink>
						<url>https://assets.yolacdn.net/store/generic/image-1.jpg</url>
					</imageLink>
				</link>
			</image>
		</product>
		<product id="0002">
			<description>This is a product description. Add more detail about this product, such as benefits, appearance, components and value</description>
			<name>Product 3</name>
			<price>19.00</price>
			<sku>0002</sku>
			<image>
				<link>
					<imageLink>
						<url>https://assets.yolacdn.net/store/generic/image-2.jpg</url>
					</imageLink>
				</link>
			</image>
		</product>
		<product id="0003">
			<description>This is a product description. Add more detail about this product, such as benefits, appearance, components and value</description>
			<name>Product 4</name>
			<price>99.00</price>
			<sku>0003</sku>
			<image>
				<link>
					<imageLink>
						<url>https://assets.yolacdn.net/store/generic/image-3.jpg</url>
					</imageLink>
				</link>
			</image>
		</product>
		<product id="0004">
			<description>This is a product description. Add more detail about this product, such as benefits, appearance, components and value</description>
			<name>Product 5</name>
			<price>49.00</price>
			<sku>0004</sku>
			<image>
				<link>
					<imageLink>
						<url>https://assets.yolacdn.net/store/generic/image-4.jpg</url>
					</imageLink>
				</link>
			</image>
		</product>
		<product id="0005">
			<description>This is a product description. Add more detail about this product, such as benefits, appearance, components and value</description>
			<name>Product 6</name>
			<price>99.00</price>
			<sku>0005</sku>
			<image>
				<link>
					<imageLink>
						<url>https://assets.yolacdn.net/store/generic/image-5.jpg</url>
					</imageLink>
				</link>
			</image>
		</product>
		<product id="0006">
			<description>This is a product description. Add more detail about this product, such as benefits, appearance, components and value</description>
			<name>Product 7</name>
			<price>99.00</price>
			<sku>0006</sku>
			<image>
				<link>
					<imageLink>
						<url>https://assets.yolacdn.net/store/generic/image-6.jpg</url>
					</imageLink>
				</link>
			</image>
		</product>
		<product id="0007">
			<description>This is a product description. Add more detail about this product, such as benefits, appearance, components and value</description>
			<name>Product 8</name>
			<price>19.00</price>
			<sku>0007</sku>
			<image>
				<link>
					<imageLink>
						<url>https://assets.yolacdn.net/store/generic/image-7.jpg</url>
					</imageLink>
				</link>
			</image>
		</product>
	</products>
	<orders/>
	<coupons/>
</template>
`
